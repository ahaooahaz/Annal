package storage

import (
	"context"

	pb "github.com/ahaooahaz/Annal/binaries/pb/gen"
	"github.com/huandu/go-sqlbuilder"
	"github.com/sirupsen/logrus"
)

func CreateWords(ctx context.Context, words []*pb.Word, db DB) (err error) {
	is := sqlbuilder.NewInsertBuilder()
	is.InsertInto(wordsTable)
	is.Cols(wordsCols[1:]...)
	for _, w := range words {
		is.Values(w.GetWord(), w.GetPronunciation(), w.GetCN(), w.GetStatus(), w.GetMeta(), w.GetCreatedAt(), w.GetUpdatedAt())
	}

	command, args := is.Build()
	_, err = db.Exec(command, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		}).Errorf("exec failed, err: %v", err.Error())
		return
	}

	return
}
