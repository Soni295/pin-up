package repository_test

import (
	"os"
	"testing"

	"github.com/Soni295/pin-up/src/model"
	"github.com/Soni295/pin-up/src/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	folder := "./tmp/"
	fileName := "save.json"
	fullName := folder + fileName
	repo := repository.NewRepository(folder, fileName)
	p := model.NewPath()

	t.Run("Check Save", func(t *testing.T) {
		if err := repo.Save(p); err != nil {
			assert.NoError(t, err)
		}
		f, err := os.ReadFile(fullName)
		if err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, f, []byte("{}"))

		if err := os.RemoveAll(folder + fileName); err != nil {
			assert.NoError(t, err)
		}

		if err := os.RemoveAll(folder); err != nil {
			assert.NoError(t, err)
		}
	})

}