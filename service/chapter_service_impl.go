package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/erikrios/my-story-dummy/entity"
	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/util/customerr"
	cfs "github.com/erikrios/my-story-dummy/util/fs"
	"github.com/erikrios/my-story-dummy/util/generator"
	"github.com/erikrios/my-story-dummy/validation"
)

type chapterServiceImpl struct {
	fs    cfs.FS
	idGen generator.IDGenerator
}

func NewChapterServiceImpl(fs cfs.FS, idGen generator.IDGenerator) *chapterServiceImpl {
	return &chapterServiceImpl{fs: fs, idGen: idGen}
}

func (s *chapterServiceImpl) Create(p payload.CreateChapter) (id string, err error) {
	if validateErr := validation.Validate.Struct(p); validateErr != nil {
		err = customerr.InvalidPayload
		return
	}

	absolutePath := fmt.Sprintf("%s/%s.json", "data", p.GroupID)

	if !s.fs.IsExists(absolutePath) {
		err = customerr.NotFound
		return
	}

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

	var isExists bool
	var index int
	for i, v := range resp.Data {
		if v.ID == p.StoryID {
			isExists = true
			index = i
			break
		}
	}

	if !isExists {
		err = customerr.NotFound
		return
	}

	id = s.idGen.GenerateChapterID()
	rand.Seed(time.Now().UnixNano())
	views := rand.Intn(10000)

	chapter := entity.Chapter{
		ID:        id,
		StoryID:   p.StoryID,
		Title:     p.Title,
		Content:   p.Content,
		Views:     uint(views),
		CreatedAt: time.Now().Format(time.RFC1123),
		UpdatedAt: time.Now().Format(time.RFC1123),
	}

	resp.Data[index].Chapters = append(resp.Data[index].Chapters, chapter)

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
