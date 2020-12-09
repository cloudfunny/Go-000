## 学习笔记
- [Week03: Go并发编程(一) goroutine](https://lailin.xyz/post/go-training-week3-goroutine.html)
## 作业

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

[homework](./Week03/../homework/main.go)

主要思路: 服务都通过 server 函数启动，server 函数 ctx 控制是否退出，只要 ctx 退出，那么这个服务也会被退出

```go
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
```
然后在 main 函数中我们复用 errgroup 的 context，因为只要 http 服务意外退出就会报错，报错时就会 cancel 掉 context，由于我们使用了同一个 context 所以只要有一个服务退出了，那么另外一个服务也会跟着退出

同时在 main 函数捕获到退出信号之后，我们就调用 cancel 函数将 context cancel 掉，这样两个服务都会优雅退出了

```go
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
	log.Println("Shutting down server...")
	cancel()

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
```