package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahaooahaz/annal/samples/golang/asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "localhost:6379"}, asynq.Config{
		Concurrency: 1,
	})

	_ = srv.Run(task.Mux)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-quitCh
	fmt.Printf("receiving signal: %v, start to quit", sig)
}
