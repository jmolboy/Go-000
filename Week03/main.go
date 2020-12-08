package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	grp, errCtx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return startHttpServer1(errCtx)
	})

	grp.Go(func() error {
		return startHttpServer2(errCtx)
	})

	grp.Go(func() error {
		return listenSig(cancelFunc)
	})

	fmt.Println("begin wait")
	if err := grp.Wait(); err != nil {
		fmt.Printf("exited: %s\n", err.Error())
	}
}

func startHttpServer1(pctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Print(res, "welcome server 1")
	})

	s := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	// 监听context
	go func() {
		<-pctx.Done()
		fmt.Println("shut down http server1")
		// 优雅关闭http server
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func startHttpServer2(pctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Print(res, "welcome server 2")
	})

	s := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	// 监听context
	go func() {
		<-pctx.Done()
		fmt.Println("shut down http server2")
		// 优雅关闭http server
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func listenSig(cancelFunc context.CancelFunc) error {
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigs:
		fmt.Println("get sig to exist")
		cancelFunc()
		// 退出
	}
	return nil
}
