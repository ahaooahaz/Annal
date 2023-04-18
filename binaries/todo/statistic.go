package todo

import (
	"context"
	"time"

	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
	"github.com/AHAOAHA/Annal/binaries/storage"
	"github.com/sirupsen/logrus"
)

type StatisticInformation struct {
	Total   int
	Done    int
	Plan    int
	Expired int
}

func Statistic(ctx context.Context) (s *StatisticInformation) {
	s = &StatisticInformation{}
	tasks, err := storage.ListTodoTasks(ctx, storage.GetInstance())
	if err != nil {
		logrus.Errorf("list todo task failed, %s", err.Error())
		return
	}
	for _, task := range tasks {
		switch task.GetStatus() {
		case pb.TodoTaskStatus_DONE:
			s.Done++
		case pb.TodoTaskStatus_PENDING:
			if time.Now().Unix() < task.GetPlan() {
				s.Expired++
			} else {
				s.Plan++
			}
		}
	}
	return
}
