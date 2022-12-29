package serve

import (
	"context"

	"github.com/AHAOAHA/Annal/binaries/internal/todo"
	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
)

type todoServe struct {
	notifies []chan *pb.TodoTask
	pb.UnimplementedTodoServiceServer
}

func (t *todoServe) CreateTodoTask(ctx context.Context, in *pb.CreateTodoTaskRequest) (out *pb.CreateTodoTaskResponse, err error) {
	if err = in.Validate(); err != nil {
		return
	}
	err = todo.CreateTodoTask(ctx, in.GetTitle(), in.GetDesp(), in.GetPlan().AsTime())
	if err != nil {
		return
	}
	return
}

func (t *todoServe) ListTodoTasks(ctx context.Context, in *pb.ListTodoTasksRequest) (out *pb.ListTodoTasksResponse, err error) {
	out = &pb.ListTodoTasksResponse{}
	out.Tasks, err = todo.ListTodoTasks(ctx)
	if err != nil {
		return
	}
	return
}
func (t *todoServe) UpdateTodoTask(ctx context.Context, in *pb.UpdateTodoTaskRequest) (out *pb.UpdateTodoTaskResponse, err error) {
	if err = in.Validate(); err != nil {
		return
	}

	return
}
func (t *todoServe) PruneTodoTasks(ctx context.Context, in *pb.PruneTodoTasksRequest) (out *pb.PruneTodoTasksResponse, err error) {
	err = todo.PruneTodoTasks(ctx)
	return
}
