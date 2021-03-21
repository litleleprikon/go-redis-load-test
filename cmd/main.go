package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/litleleprikon/go-redis-load-test/internal/reporter"
	"github.com/litleleprikon/go-redis-load-test/internal/runner"
)

var durationFlag = flag.Duration("duration", time.Minute, "duration of test")
var configFlag = flag.String("config", "./config.yml", "path to config file")

func main() {
	flag.Parse()
	ctx := context.Background()
	var done context.CancelFunc
	if durationFlag != nil {
		ctx, done = context.WithTimeout(ctx, *durationFlag)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go shutdown(sigs, done)

	var executor runner.Executor
	var rep reporter.Reporter
	runner := runner.NewRunner(executor, rep)
	if err := runner.Run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func shutdown(sigs chan os.Signal, done context.CancelFunc) {
	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)
	done()
}
