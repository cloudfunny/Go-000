## 学习笔记
- [Week03: Go并发编程(一) goroutine](https://lailin.xyz/post/go-training-week3-goroutine.html)
## 作业

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

[homework](./Week03/homework/main.go)

```go
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 模拟单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// g1
	// g1 退出了所有的协程都能退出么？
	// g1 退出后, context 将不再阻塞，g2, g3 都会随之退出
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// g2
	// g2 退出了所有的协程都能退出么？
	// g2 退出时，调用了 shutdown，g1 会退出
	// g2 退出后, context 将不再阻塞，g3 会随之退出
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverOut:
			log.Println("server will out...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	// g3
	// g3 捕获到 os 退出信号将会退出
	// g3 退出了所有的协程都能退出么？
	// g3 退出后, context 将不再阻塞，g2 会随之退出
	// g2 退出时，调用了 shutdown，g1 会退出
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
```

## 评价记录
上版本助教老师的评语:

对errgroup的理解都稍微有一点问题。

1.使用errgroup就不需要再自己cancel了

me: 这个不是理解问题，这个在第一个版本是有意这么用的，这样可以在捕获到信号之后优雅退出两个 http server

2.你的信号处理代码没有作为一个单独的goroutine放到errgroup中，如果ListenAndServe出错异常退出，你的进程将不能正常退出;

me: 这个的确是有问题的, 没有考虑到

3.你的Shutdown协程没有放到errgroup中，这会导致errgroup.Wait不会等它，最终导致已经连接的请求可能无法在进程退出前合适得处理完

me: 偷懒了，翻了一下源码发现，如果没有超时控制会导致shutdown立即退出