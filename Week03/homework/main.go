package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)

	g.Go(func() error {
		return server(ctx, http.DefaultServeMux, ":8081")
	})

	g.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		return server(ctx, mux, ":8080")
	})

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("")
	log.Println("shutting down server...")
	cancel()

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

// 启动一个服务
func server(ctx context.Context, handler http.Handler, addr string) error {
	s := http.Server{
		Handler: handler,
		Addr:    addr,
	}

	// 这个 goroutine 我们可以控制退出，因为只要 stop 这个 channel close 或者是写入数据，这里就会退出
	// 同时因为调用了 s.Shutdown 调用之后，http 这个函数启动的 http server 也会优雅退出
	go func() {
		<-ctx.Done()
		log.Printf("server will exiting, addr: %s", addr)
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
