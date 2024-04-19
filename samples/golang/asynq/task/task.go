package task

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	HANDLE string = "handle"
)

var (
	Mux *asynq.ServeMux
)

func init() {
	Mux = asynq.NewServeMux()
	Mux.HandleFunc(HANDLE, Handle)
}

func Handle(ctx context.Context, t *asynq.Task) (err error) {
	taskID, _ := asynq.GetTaskID(ctx)
	fmt.Println(taskID, "PROCESSING")
	select {
	case <-ctx.Done():
		fmt.Println(taskID, "CANCEL")
	case <-time.After(time.Second * 30):
		fmt.Println(taskID, "FINISH")
	}
	return
}
