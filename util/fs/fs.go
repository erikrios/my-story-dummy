package fs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type FS interface {
	CreateDir(name string) error
	CreateFile(name string) error
	WriteFile(name string, data []byte) error
	ReadFile(name string) ([]byte, error)
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

// IsExists check whether a directory with given name is exists relatively to the working directory
func (l *localFS) IsExists(name string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", l.wd, name))
	return !os.IsNotExist(err)
}

// CreateFile create a file with given name relatively to the working directory
func (l *localFS) CreateFile(name string) error {
	_, err := os.Create(fmt.Sprintf("%s/%s", l.wd, name))
	return err
}

// WriteFile write bytes of data to given file name relatively to the working directory
func (l *localFS) WriteFile(name string, data []byte) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", l.wd, name), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// ReadFile read given file name to bytes of data relatively to the working directory
func (l *localFS) ReadFile(name string) ([]byte, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s", l.wd, name))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buffer := make([]byte, 1024)
	data := make([]byte, 0, 1024)

	for {
		read, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		data = append(data, buffer[:read]...)
	}

	return data, nil
}
