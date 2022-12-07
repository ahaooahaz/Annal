package storage

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/sirupsen/logrus"
)

func ListTodoTasks(ctx context.Context) (tasks []*TodoTask, err error) {
	ss := sqlbuilder.NewSelectBuilder()
	ss.From(todosTable)
	ss.Select(todosCols...)

	command, args := ss.Build()

	err = getInstance().Select(&tasks, command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func CreateTodoTasks(ctx context.Context, tasks []*TodoTask) (err error) {
	ss := sqlbuilder.NewInsertBuilder()
	ss.InsertInto(todosTable)
	ss.Cols(todosCols[1:]...)
	for _, task := range tasks {
		ss.Values(task.UUID, task.Title, task.Description, task.Plan, task.Status, task.CreatedAt, task.UpdatedAt)
	}
	command, args := ss.Build()
	_, err = getInstance().Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func DeleteTodoTasks(ctx context.Context, ids []interface{}) (err error) {
	sd := sqlbuilder.NewDeleteBuilder()
	sd.DeleteFrom(todosTable)
	sd.Where(
		sd.In(todosCols[1], ids...),
	)
	command, args := sd.Build()
	_, err = getInstance().Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
	}
	return
}

func UpdateTodoTask(ctx context.Context, task *TodoTask) (err error) {
	su := sqlbuilder.NewUpdateBuilder()
	su.Update(todosTable)
	su.Set(
		su.E(todosCols[3], task.Title),
		su.E(todosCols[4], task.Description),
		su.E(todosCols[5], task.Plan),
		su.E(todosCols[6], task.Status),
		su.E(todosCols[8], task.UpdatedAt),
	)
	su.Where(
		su.E(todosCols[1], task.UUID),
	)
	command, args := su.Build()
	_, err = getInstance().Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		})
	}
	return
}
