//go:build gocv
// +build gocv

package video

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateVideoFromImage(t *testing.T) {
	err := GenerateVideoFromImage([]string{"image.jpg"}, "test.avi", 25, 512, 512, 1000)
	defer os.Remove("test.avi")
	assert.Equal(t, nil, err)
}
