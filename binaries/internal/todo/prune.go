package todo

import (
	"context"

	proto "github.com/AHAOAHA/Annal/binaries/internal/pb/gen"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "prune task",
	Long:  `prune task`,
	Run:   pruneTodoTask,
}

func pruneTodoTask(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	var tx *sqlx.Tx
	tx, err = storage.GetInstance().Beginx()
	if err != nil {
		logrus.Errorf("%v", err.Error())
		return
	}

	var tasks []*proto.TodoTask
	tasks, err = storage.ListTodoTasksWithCondition(ctx, tx, map[string]interface{}{
		"status": proto.TodoTaskStatus_DONE.Number(),
	})
	if err != nil {
		logrus.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}

	ids := []interface{}{}
	for _, id := range tasks {
		ids = append(ids, id.GetUUID())
	}

	err = storage.DeleteTodoTasks(ctx, tx, ids)
	if err != nil {
		logrus.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		logrus.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}
}
