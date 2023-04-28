package todo

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ahaooahaz/Annal/binaries/config"
	"github.com/ahaooahaz/Annal/binaries/notify"
	pb "github.com/ahaooahaz/Annal/binaries/pb/gen"
	"github.com/ahaooahaz/Annal/binaries/storage"
	"github.com/ahaooahaz/encapsutils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	createCmd.Flags().StringP("title", "t", "", "todotosk title (max length 256)")
	createCmd.Flags().StringP("desp", "d", "", "todotosk desp (max length 1024)")
	createCmd.Flags().Uint64P("notify", "n", 3, "notify timeout seconds")
	createCmd.Flags().StringP("plan", "p", time.Now().Add(time.Hour).Format(_TimeFormatString), fmt.Sprintf("plan time, format: %s", _TimeFormatString))

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("desp")
	createCmd.MarkFlagRequired("plan")
}

func createTodoTask(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error

	fetch(ctx)

	var title, desp string

	title, err = cmd.Flags().GetString("title")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	if strings.TrimSpace(title) == "" {
		fmt.Fprintf(os.Stderr, "title format invalid")
		return
	}

	desp, err = cmd.Flags().GetString("desp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	var planStr string
	planStr, err = cmd.Flags().GetString("plan")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	var plan time.Time
	plan, err = time.ParseInLocation(_TimeFormatString, planStr, time.Local)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse plan time failed, err: %v", err.Error())
		return
	}

	var notifyTimeout uint64
	notifyTimeout, err = cmd.Flags().GetUint64("notify")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	err = CreateTodoTask(ctx, title, desp, plan, notifyTimeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}

func CreateTodoTask(ctx context.Context, title, desp string, plan time.Time, notifyTimeout uint64) (err error) {
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

	commandLine := []string{"#!/bin/bash"}
	jobpath := config.ATJOBS + "/" + task.GetUUID() + ".sh"
	step1 := []string{config.NOTIFYSENDSH, "-ti", fmt.Sprintf("'%s'", task.GetTitle()), "-d", fmt.Sprintf("'%s'", task.GetDescription()), "-t", fmt.Sprintf("%d", 10)}
	step2 := []string{"rm", "-rf", jobpath}
	strings.Join(step1, " ")
	commandLine = append(commandLine, strings.Join(step1, " "))
	commandLine = append(commandLine, strings.Join(step2, " "))

	var f *os.File
	f, err = encapsutils.CreateFile(jobpath)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	defer f.Close()

	write := bufio.NewWriter(f)
	for _, c := range commandLine {
		_, err = write.WriteString(c)
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
		_, err = write.WriteString("\n")
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
	}
	write.Flush()

	var jobID uint64
	jobID, err = notify.CreateOnTimeJob(ctx, jobpath, time.Unix(task.GetPlan(), 0))
	if err != nil {
		return
	}

	task.NotifyJobId = jobID
	err = storage.CreateTodoTasks(ctx, storage.GetInstance(), []*pb.TodoTask{task})
	if err != nil {
		return
	}
	return
}
