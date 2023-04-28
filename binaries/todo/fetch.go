package todo

import (
	"context"
	"sync"
	"time"

	"github.com/ahaooahaz/Annal/binaries/pb/gen"
	"github.com/ahaooahaz/Annal/binaries/storage"
	"github.com/sirupsen/logrus"
)

func fetch(ctx context.Context) {
	tasks, err := storage.ListTodoTasks(ctx, storage.GetInstance())
	if err != nil {
		logrus.Errorf("list todo tasks failed, err: %v", err.Error())
		return
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t *gen.TodoTask) {
			defer wg.Done()
			if t.GetStatus() != gen.TodoTaskStatus_DONE {
				if time.Unix(t.GetPlan(), 0).Before(time.Now()) {
					t.Status = gen.TodoTaskStatus_EXPIRED
				}

				ine := storage.UpdateTodoTask(ctx, storage.GetInstance(), t)
				if ine != nil {
					logrus.Errorf("update todo task failed, err: %v", ine.Error())
					return
				}
			}
		}(task)
	}
	wg.Wait()
}
