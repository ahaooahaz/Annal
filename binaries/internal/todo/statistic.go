package todo

import (
	"context"
	"time"

	proto "github.com/AHAOAHA/Annal/binaries/internal/pb/gen"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/sirupsen/logrus"
)

type StatisticInformation struct {
	Total   int
	Done    int
	Plan    int
	Expired int
}

func Statistic(ctx context.Context) (s *StatisticInformation) {
	tasks, err := storage.ListTodoTasks(ctx, storage.GetInstance())
	if err != nil {
		logrus.Errorf("list todo task failed, %s", err.Error())
		return
	}
	for _, task := range tasks {
		switch task.GetStatus() {
		case proto.TodoTaskStatus_DONE:
			s.Done++
		case proto.TodoTaskStatus_PENDING:
			if time.Now().Unix() < task.GetPlan() {
				s.Expired++
			} else {
				s.Plan++
			}
		}
	}
	return
}
