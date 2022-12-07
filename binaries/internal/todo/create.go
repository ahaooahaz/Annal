package todo

import (
	"context"
	"fmt"
	"time"

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
	task := &storage.TodoTask{
		UUID:      uuid.New().String(),
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	fmt.Print("title:")
	fmt.Scanf("%s", &task.Title)
	fmt.Print("desp:")
	fmt.Scanf("%s", &task.Description)

	err := storage.CreateTodoTasks(ctx, []*storage.TodoTask{task})
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
}
