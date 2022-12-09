package todo

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	proto "github.com/AHAOAHA/Annal/binaries/internal/pb/gen"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
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

func createTodoTask(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	task := &proto.TodoTask{
		UUID:      uuid.New().String(),
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What todo:\n")
	title, _, _ := reader.ReadLine()
	task.Title = string(title)
	fmt.Printf("What desp:\n")
	desp, _, _ := reader.ReadLine()
	task.Description = string(desp)
	fmt.Printf("When todo: (default: 1 hour later)\n")
	var planT int
	var plan time.Time

	fmt.Scanf("%d\n", planT)
	switch planT {
	case 1:
		fmt.Printf("Input time: (format: %s)\n", _TimeFormatString)
		planS, _, _ := reader.ReadLine()
		plan, err = time.Parse(_TimeFormatString, string(planS))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}
	default:
		plan = time.Now().Add(time.Hour)
	}
	task.Plan = plan.Unix()

	err = storage.CreateTodoTasks(ctx, storage.GetInstance(), []*proto.TodoTask{task})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		return
	}
}
