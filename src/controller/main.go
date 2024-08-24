package controller

import (
	"fmt"

	"github.com/Soni295/pin-up/src/model"
	"github.com/Soni295/pin-up/src/repository"
)

type MoveAction interface {
	Move(string) error
}
type MockMover struct {
	path string
}

func (mm *MockMover) Move(path string) error {
	mm.path = path
	return nil
}

type PathCtrl struct {
	path *model.Path
	repo *repository.Repository
	move MoveAction
}

func NewPathCtrl(path *model.Path, repo *repository.Repository) *PathCtrl {
	return &PathCtrl{
		path: path,
		repo: repo,
	}
}

func (pc *PathCtrl) Save(key string, path string) error {
	pc.path.Add(key, path)
	return pc.repo.Save(pc.path)
}

func (pc *PathCtrl) Move(key string) error {
	path, err := pc.repo.Get()
	if err != nil {
		return err
	}
	direction := path.Get(key)
	if direction == "" {
		return fmt.Errorf("No path save.")
	}
	return pc.move.Move(direction)
}