package fs

import (
	"fmt"
	"log"
	"os"
)

type FS interface {
	CreateDir(name string) error
	CreateFile(name string) error
	IsExists(name string) bool
}

type localFS struct {
	wd string
}

func NewLocalFS() *localFS {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return &localFS{wd: wd}
}

// CreateDir create a directory with given name relatively to the working directory
func (l *localFS) CreateDir(name string) error {
	return os.Mkdir(fmt.Sprintf("%s/%s", l.wd, name), os.ModePerm)
}

func (l *localFS) IsExists(name string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", l.wd, name))
	return !os.IsNotExist(err)
}

func (l *localFS) CreateFile(name string) error {
	_, err := os.Create(fmt.Sprintf("%s/%s", l.wd, name))
	return err
}
