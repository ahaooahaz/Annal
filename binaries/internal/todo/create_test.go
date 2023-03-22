package todo

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoTask(t *testing.T) {
	config.ATJOBS = "../../../.at.jobs/"
	config.DBPATH = "test.db"
	config.NOTIFYSENDSH = "notify-send.sh"
	err := storage.Sqlite3Migrate("../../../migrations/sqlite3", config.DBPATH)
	assert.Equal(t, nil, err)
	defer os.Remove(config.DBPATH)
	err = CreateTodoTask(context.Background(), "test", "test", time.Now(), 1)
	// err = CreateTodoTask(context.Background(), "test", "test", time.Now(), notify)
	assert.Equal(t, nil, err)
}
