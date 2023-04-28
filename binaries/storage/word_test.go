package storage

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/ahaooahaz/Annal/binaries/pb/gen"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateWords(t *testing.T) {
	dbx, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Equal(t, nil, err)
	defer dbx.Close()

	dbv := sqlx.NewDb(dbx, "")

	words := []*pb.Word{
		{
			Word:          "test",
			Pronunciation: "test",
			CN:            "test",
			Status:        1,
			Meta:          "test",
			CreatedAt:     1,
			UpdatedAt:     1,
		},
		{
			Word:          "test",
			Pronunciation: "test",
			CN:            "test",
			Status:        1,
			Meta:          "test",
			CreatedAt:     1,
			UpdatedAt:     1,
		},
	}

	// expect insert single words object.
	EXPECT_InsertSQL(mock, wordsTable, wordsCols[1:], []interface{}{
		words[0],
	}, 1)
	err = CreateWords(context.Background(), []*pb.Word{words[0]}, dbv)
	assert.Equal(t, nil, err)

	EXPECT_InsertSQL(mock, wordsTable, wordsCols[1:], []interface{}{
		words[0],
		words[1],
	}, 1)
	err = CreateWords(context.Background(), words, dbv)
	assert.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	assert.Equal(t, nil, err)

}
