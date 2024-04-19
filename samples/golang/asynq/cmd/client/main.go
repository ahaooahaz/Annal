package main

import (
	"flag"
	"fmt"

	"github.com/ahaooahaz/annal/samples/golang/asynq/task"
	"github.com/hibiken/asynq"
)

var (
	cancel = flag.Bool("cancel", false, "cancel task")
	delete = flag.Bool("delete", false, "delete task")
	taskID = flag.String("task_id", "", "task_id")

	qu = "default"
)

func main() {
	flag.Parse()
	cli := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	ipt := asynq.NewInspector(asynq.RedisClientOpt{Addr: "localhost:6379"})

	if *cancel {
		err := ipt.CancelProcessing(*taskID)
		if err != nil {
			panic(err)
		}
		return
	} else if *delete {
		err := ipt.DeleteTask(qu, *taskID)
		if err != nil {
			panic(err)
		}
		return
	}

	info, err := cli.Enqueue(asynq.NewTask(task.HANDLE, []byte("hello")), asynq.TaskID(*taskID))
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
