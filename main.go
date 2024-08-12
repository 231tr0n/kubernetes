package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"sync"

	// "net/http/httputil"
	_ "crypto/tls/fipsonly"
	_ "net/http/pprof"

	// "net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type wrapperResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapperResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

var mu *sync.Mutex

func init() {
	mu = &sync.Mutex{}
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	defer func() {
		if err := recover(); err != nil {
			err, ok := err.(error)
			if ok {
				slog.Error(err.Error())
				os.Exit(1)
			}
			log.Fatalln("FATAL:", err)
		}
	}()

	port, exists := os.LookupEnv("PORT")
	if !exists {
		slog.Error("PORT environmental variable not set")
		os.Exit(1)
	}

	// pingServiceHost, exists := os.LookupEnv("PING_HOST")
	// if !exists {
	// 	slog.Error("PING service host not set")
	// 	os.Exit(1)
	// }
	//
	// pongServiceHost, exists := os.LookupEnv("PONG_HOST")
	// if !exists {
	// 	slog.Error("PONG service host not set")
	// 	os.Exit(1)
	// }
	//
	// pingServiceUrl, err := url.Parse(pingServiceHost)
	// if err != nil {
	// 	panic(err)
	// }
	// pongServiceUrl, err := url.Parse(pongServiceHost)
	// if err != nil {
	// 	panic(err)
	// }

	// http.Handle("/ping/", httputil.NewSingleHostReverseProxy(pingServiceUrl))
	// http.Handle("/pong/", httputil.NewSingleHostReverseProxy(pongServiceUrl))
	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Service working"))
		if err != nil {
			slog.Error(err.Error())
		}
	})
	http.HandleFunc("POST /test/{value}", func(w http.ResponseWriter, r *http.Request) {
		val := r.PathValue("value")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Service working: %s", val)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 50 << 20,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			startTime := time.Now()
			wr := &wrapperResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusAccepted,
			}
			time.Sleep(time.Second)
			http.DefaultServeMux.ServeHTTP(wr, r)
			if wr.statusCode > 399 {
				slog.Warn("REQUEST", "method", r.Method, "ip", r.RemoteAddr, "time", time.Since(startTime).Round(time.Second), "status", wr.statusCode, "path", r.URL.Path)
				mu.Unlock()
				return
			}
			slog.Info("REQUEST", "method", r.Method, "ip", r.RemoteAddr, "time", time.Since(startTime).Round(time.Second), "status", wr.statusCode, "path", r.URL.Path)
			mu.Unlock()
		}),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
			slog.Warn(err.Error())
		}
	}()
	slog.Info("Started server", "port", port)

	interrupt := make(chan os.Signal, 1)
	defer close(interrupt)

	go (func() {
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	})()

	val, ok := <-interrupt
	if !ok {
		slog.Warn("Channel closed before receiving os signal")
	}
	log.Println()
	slog.Warn(val.String() + " received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Shutdown server", "port", port)
}
