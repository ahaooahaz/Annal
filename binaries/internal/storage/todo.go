package storage

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/sirupsen/logrus"

	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
)

func ListTodoTasks(ctx context.Context, db DB) (tasks []*pb.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(TodosCols...)

	command, args := ss.Build()

	err = db.Select(&tasks, command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func ListTodoTasksWithCondition(ctx context.Context, db DB, conditions map[string]interface{}) (tasks []*pb.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(TodosCols...)
	conditionsS := []string{}
	for k, v := range conditions {
		conditionsS = append(conditionsS, ss.E(k, v))
	}
	ss.Where(
		conditionsS...,
	)
	command, args := ss.Build()

	err = db.Select(&tasks, command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func CreateTodoTasks(ctx context.Context, db DB, tasks []*pb.TodoTask) (err error) {
	ss := sqlbuilder.NewInsertBuilder()
	ss.InsertInto(todosTable)
	ss.Cols(TodosCols[1:]...)
	for _, task := range tasks {
		ss.Values(task.UUID, task.Index, task.Title, task.Description, task.Plan, task.Status, task.CreatedAt, task.UpdatedAt)
	}
	command, args := ss.Build()
	_, err = db.Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func DeleteTodoTasks(ctx context.Context, db DB, ids []interface{}) (err error) {
	sd := sqlbuilder.NewDeleteBuilder()
	sd.DeleteFrom(todosTable)
	sd.Where(
		sd.In(TodosCols[1], ids...),
	)
	command, args := sd.Build()
	_, err = db.Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func UpdateTodoTask(ctx context.Context, db DB, task *pb.TodoTask) (err error) {
	su := sqlbuilder.NewUpdateBuilder()
	su.Update(todosTable)

	if task.Status == pb.TodoTaskStatus_DONE {
		su.Set(
			su.E(TodosCols[2], 0),
			su.E(TodosCols[3], task.Title),
			su.E(TodosCols[4], task.Description),
			su.E(TodosCols[5], task.Plan),
			su.E(TodosCols[6], task.Status),
			su.E(TodosCols[8], task.UpdatedAt),
		)
	} else {
		su.Set(
			su.E(TodosCols[3], task.Title),
			su.E(TodosCols[4], task.Description),
			su.E(TodosCols[5], task.Plan),
			su.E(TodosCols[6], task.Status),
			su.E(TodosCols[8], task.UpdatedAt),
		)
	}
	su.Where(
		su.E(TodosCols[2], task.Index),
	)

	command, args := su.Build()
	_, err = db.Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		})
	}
	return
}

func SelectTodoTask(ctx context.Context, db DB, ID int64) (task *pb.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(TodosCols...)
	ss.Where(
		ss.E(TodosCols[0], ID),
	)
	ss.Limit(1)

	command, args := ss.Build()
	task = &pb.TodoTask{}
	err = db.Get(task, command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}

	return
}
