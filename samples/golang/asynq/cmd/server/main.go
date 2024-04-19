package main

import (
	"github.com/ahaooahaz/annal/samples/golang/asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "localhost:6379"}, asynq.Config{
		Concurrency: 1,
	})

	_ = srv.Run(task.Mux)
}
