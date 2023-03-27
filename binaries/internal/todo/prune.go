package todo

import (
	"context"
	"fmt"
	"os"

	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/AHAOAHA/Annal/binaries/internal/utils"
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

	fetch(ctx)

	err := PruneTodoTasks(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}

func PruneTodoTasks(ctx context.Context) (err error) {
	log := logrus.WithFields(logrus.Fields{
		"act": "Prune",
	})
	var tx *sqlx.Tx
	tx, err = storage.GetInstance().Beginx()
	if err != nil {
		log.Errorf("%v", err.Error())
		return
	}

	var tasks []*pb.TodoTask
	tasks, err = storage.ListTodoTasksWithCondition(ctx, tx, false, []*utils.Pair{
		{K: "status", V: pb.TodoTaskStatus_DONE.Number()}, {K: "status", V: pb.TodoTaskStatus_EXPIRED.Number()},
	})
	log = log.WithField("tasks", tasks)
	if err != nil {
		log.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}

	ids := []interface{}{}
	for _, id := range tasks {
		ids = append(ids, id.GetUUID())
	}

	err = storage.DeleteTodoTasks(ctx, tx, ids)
	if err != nil {
		log.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		log.Errorf("%v", err.Error())
		tx.Rollback()
		return
	}

	log.Infof("done")
	return
}
