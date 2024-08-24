package model_test

import (
	"fmt"
	"testing"

	"github.com/Soni295/pin-up/src/model"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Run("Check Create", func(t *testing.T) {
		path := model.NewPath()
		assert.NotNil(t, path)
		key := "example"
		pathName := "$HOME"

		t.Run("Check Get Path that doesn't exist", func(t *testing.T) {
			output := path.Get(key)
			assert.Equal(t, output, "")
		})

		t.Run("Check Get Path", func(t *testing.T) {
			path.Add(key, pathName)
			output := path.Get(key)
			assert.Equal(t, output, pathName)
		})

		t.Run("Remove Path", func(t *testing.T) {
			path.Remove(key)
			output := path.Get(key)
			assert.Equal(t, output, "")
		})
	})

	t.Run("Check Write", func(t *testing.T) {
		t.Run("Successfully", func(t *testing.T) {
			path := model.NewPath()
			key := "home"
			pathName := "$HOME"
			info := fmt.Sprintf(`{%q:%q}`, key, pathName)
			_, err := path.Write([]byte(info))
			assert.NoError(t, err)
			output := path.Get(key)
			assert.Equal(t, pathName, output)
		})

		t.Run("Failed", func(t *testing.T) {
			path := model.NewPath()
			info := ""
			_, err := path.Write([]byte(info))
			assert.Error(t, err)
		})
	})

}