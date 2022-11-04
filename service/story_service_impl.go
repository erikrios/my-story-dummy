package service

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/erikrios/my-story-dummy/entity"
	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/util/customerr"
	cfs "github.com/erikrios/my-story-dummy/util/fs"
	"github.com/erikrios/my-story-dummy/util/generator"
	"github.com/erikrios/my-story-dummy/validation"
)

type storyServiceImpl struct {
	fs    cfs.FS
	idGen generator.IDGenerator
}

func NewStoryServiceImpl(fs cfs.FS, idGen generator.IDGenerator) *storyServiceImpl {
	return &storyServiceImpl{fs: fs, idGen: idGen}
}

func (s *storyServiceImpl) Create(p payload.CreateStory) (id string, err error) {
	if validateErr := validation.Validate.Struct(p); validateErr != nil {
		err = customerr.InvalidPayload
		return
	}

	absolutePath := fmt.Sprintf("%s/%s.json", "data", p.GroupID)

	if !s.fs.IsExists(absolutePath) {
		err = customerr.NotFound
		return
	}

	wd, wdErr := os.Getwd()
	if wdErr != nil {
		err = customerr.Internal
		return
	}

	posterPath := fmt.Sprintf("%s/assets/%s", wd, p.Poster)

	_, statErr := os.Stat(posterPath)
	if os.IsNotExist(statErr) {
		err = customerr.NotFound
		return
	}

	f, openErr := os.Open(posterPath)
	if openErr != nil {
		err = customerr.Internal
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buffer := make([]byte, 1024)
	posterBytes := make([]byte, 0, 1024)

	for {
		read, readErr := reader.Read(buffer)
		if readErr != nil {
			if readErr != io.EOF {
				err = customerr.Internal
				return
			}
			break
		}
		posterBytes = append(posterBytes, buffer[:read]...)
	}

	poster := base64.StdEncoding.EncodeToString(posterBytes)

	data, readErr := s.fs.ReadFile(absolutePath)
	if readErr != nil {
		err = customerr.Internal
		return
	}

	var resp entity.Response

	if unmarshalErr := json.Unmarshal(data, &resp); unmarshalErr != nil {
		err = customerr.Internal
		return
	}

	id = s.idGen.GenerateStoryID()

	story := entity.Story{
		ID:        id,
		Title:     p.Title,
		Poster:    poster,
		CreatedAt: time.Now().Format(time.RFC1123),
		UpdatedAt: time.Now().Format(time.RFC1123),
		Chapters:  []entity.Chapter{},
	}

	resp.Data = append(resp.Data, story)

	respBytes, marshalErr := json.MarshalIndent(resp, "", "  ")
	if marshalErr != nil {
		err = customerr.Internal
		return
	}

	if writeErr := s.fs.WriteFile(absolutePath, respBytes); writeErr != nil {
		err = customerr.Internal
		return
	}

	return
}
