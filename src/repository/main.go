package repository

import (
	"io"
	"os"

	"github.com/Soni295/pin-up/src/model"
)

type Repository struct {
	pathFile string
	fileName string
}

const defaultFilePath = "$HOME/.example/"
const defaultFileName = "repo.json"

func NewRepository(path string, fileName string) *Repository {
	if path == "" {
		path = defaultFilePath
	}
	if fileName == "" {
		fileName = defaultFileName
	}
	return &Repository{
		pathFile: path,
		fileName: fileName,
	}
}

func (r *Repository) Save(info io.Reader) error {
	_, err := os.Stat(r.pathFile)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(r.pathFile, os.ModePerm); err != nil {
			return err
		}
	}
	file, err := os.Create(r.pathFile + r.fileName)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(info)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}

func (r *Repository) Get() (*model.Path, error) {
	bytes, err := os.ReadFile(r.pathFile)
	if err != nil {
		return nil, err
	}
	p := model.NewPath()

	_, err = p.Write(bytes)
	if err != nil {
		return nil, err
	}

	return p, nil
}