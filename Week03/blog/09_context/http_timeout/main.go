package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

// func rpc(ctx context.Context, ms int) (string, error) {
// 	select {
// 	default:
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	}
// 	time.Sleep(time.Duration(ms) * time.Millisecond)
// 	return "rpc done", nil
// }

// func handle(ctx context.Context) {
// 	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
// 	defer cancel()
// 	for _, v := range v {
// 		rpc(ctx)
// 	}
// 	res, err := rpc(ctx)
// 	fmt.Println(res)
// 	fmt.Println(err)
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// 这里阻塞住，goroutine 不会释放的
		// time.Sleep(1000 * time.Second)
	})
	handler := http.TimeoutHandler(mux, time.Millisecond, "xxx")
	go func() {
		if err := http.ListenAndServe("0.0.0.0:8066", nil); err != nil {
			panic(err)
		}
	}()
	http.ListenAndServe(":8080", handler)

	// for {
	// 	time.Sleep(1 * time.Second)
	// 	go handle(context.Background())
	// }
}
