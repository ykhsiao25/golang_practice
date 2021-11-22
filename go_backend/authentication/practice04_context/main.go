package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// /bar 流程:call bar() -> initialize ctx.Value() -> dbAccess() to get the result
// -> ctx set timeout -> goroutine get the value -> goroutine write the value to chan
// -> if not timeout -> return the value to bar() -> write to res
// -> if  timeout -> return 0 and ctx.Err() to bar() -> write to res
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("./favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userid", 777)
	ctx = context.WithValue(ctx, "fname", "bond")

	result, err := dbAccess(ctx)
	if err != nil {
		http.Error(res, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(res, result)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//確保ctx都進行回收
	defer cancel()

	ch := make(chan int)

	//這邊事實上在main routine執行完才執行 (因為10秒太久 上面是設定一秒)
	go func() {
		uid := ctx.Value("userid").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			return
		}
		ch <- uid
	}()

	//這邊是main routine 該執行的地方(若main routime等不到 goroutine的ch value，就會到 case <- ctx.Done())
	select {
	//收到cancel() ，執行這裡
	case <-ctx.Done():
		return 0, ctx.Err()

	case i := <-ch:
		return i, nil
	}
}
