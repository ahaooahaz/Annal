package storage

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/sirupsen/logrus"

	proto "github.com/AHAOAHA/Annal/binaries/internal/pb/gen"
)

func ListTodoTasks(ctx context.Context, db DB) (tasks []*proto.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(todosCols...)

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

func ListTodoTasksWithCondition(ctx context.Context, db DB, conditions map[string]interface{}) (tasks []*proto.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(todosCols...)
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

func CreateTodoTasks(ctx context.Context, db DB, tasks []*proto.TodoTask) (err error) {
	ss := sqlbuilder.NewInsertBuilder()
	ss.InsertInto(todosTable)
	ss.Cols(todosCols[1:]...)
	for _, task := range tasks {
		ss.Values(task.UUID, task.Title, task.Description, task.Plan, task.Status, task.CreatedAt, task.UpdatedAt)
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
		sd.In(todosCols[1], ids...),
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

func UpdateTodoTask(ctx context.Context, db DB, task *proto.TodoTask) (err error) {
	su := sqlbuilder.NewUpdateBuilder()
	su.Update(todosTable)
	su.Set(
		su.E(todosCols[2], task.Title),
		su.E(todosCols[3], task.Description),
		su.E(todosCols[4], task.Plan),
		su.E(todosCols[5], task.Status),
		su.E(todosCols[7], task.UpdatedAt),
	)
	su.Where(
		su.E(todosCols[1], task.UUID),
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

func SelectTodoTask(ctx context.Context, db DB, ID int64) (task *proto.TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(todosCols...)
	ss.Where(
		ss.E(todosCols[0], ID),
	)
	ss.Limit(1)

	command, args := ss.Build()
	task = &proto.TodoTask{}
	err = db.Get(task, command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}

	return
}
