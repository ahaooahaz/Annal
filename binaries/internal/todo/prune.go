package todo

import (
	"context"
	"fmt"
	"os"

	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "prune task",
	Long:  `prune task`,
	Run:   pruneTodoTasks,
}

func pruneTodoTasks(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := PruneTodoTasks(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}

func PruneTodoTasks(ctx context.Context) (err error) {
	var tx *sqlx.Tx
	tx, err = storage.GetInstance().Beginx()
	if err != nil {
		logrus.Errorf("%v", err.Error())
		return
	}

	var tasks []*pb.TodoTask
	tasks, err = storage.ListTodoTasksWithCondition(ctx, tx, map[string]interface{}{
		"status": pb.TodoTaskStatus_DONE.Number(),
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
	return
}
