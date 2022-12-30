package todo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "create task",
	Long:    `create task`,
	Run:     createTodoTask,
}

func init() {
	createCmd.Flags().StringP("title", "t", "", "todotosk title")
	createCmd.Flags().StringP("desp", "d", "", "todotosk desp")
	createCmd.Flags().Int64P("plan", "p", 0, "plan time")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("desp")
	createCmd.MarkFlagRequired("plan")
}

func createTodoTask(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error

	var title, desp string

	title, err = cmd.Flags().GetString("title")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	desp, err = cmd.Flags().GetString("desp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	var plan int64
	plan, err = cmd.Flags().GetInt64("plan")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	err = CreateTodoTask(ctx, title, desp, time.Unix(plan, 0), notify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}

func CreateTodoTask(ctx context.Context, title, desp string, plan time.Time, notify func(task *pb.TodoTask) error) (err error) {
	task := &pb.TodoTask{
		UUID:        uuid.NewString(),
		Title:       title,
		Description: desp,
		Plan:        plan.Unix(),
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	var tasks []*pb.TodoTask
	tasks, err = storage.ListTodoTasks(ctx, storage.GetInstance())
	if err != nil {
		return
	}

	useIndexs := make(map[int64]bool)
	for _, t := range tasks {
		useIndexs[t.GetIndex()] = true
	}

	var index int64
	for index = 1; index <= 100; index++ {
		_, used := useIndexs[index]
		if !used {
			break
		}
	}
	task.Index = index

	err = storage.CreateTodoTasks(ctx, storage.GetInstance(), []*pb.TodoTask{task})
	if err != nil {
		return
	}

	err = notify(task)
	if err != nil {
		return
	}
	return
}
