package model

import (
	"encoding/json"
	"io"
)

type Path map[string]string

func NewPath() *Path { return &Path{} }

func (p Path) Add(key string, path string) { p[key] = path }
func (p Path) Get(key string) string       { return p[key] }
func (p Path) Remove(key string)           { delete(p, key) }

func (p *Path) String() string {
	j, _ := json.Marshal(p)
	return string(j)
}

func (p *Path) Read(b []byte) (int, error) {
	j, err := json.Marshal(p)
	if err != nil {
		return 1, err
	}
	return copy(b, j), io.EOF
}

func (p *Path) Write(b []byte) (int, error) {
	if err := json.Unmarshal(b, p); err != nil {
		return 1, err
	}
	return 0, nil
}