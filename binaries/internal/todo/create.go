package todo

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What todo:\n")
	title, _, _ := reader.ReadLine()
	task.Title = string(title)
	fmt.Print("What desp:\n")
	desp, _, _ := reader.ReadLine()
	task.Description = string(desp)

	err := storage.CreateTodoTasks(ctx, []*storage.TodoTask{task})
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
}
