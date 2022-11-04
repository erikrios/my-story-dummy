package service

import "github.com/erikrios/my-story-dummy/model/payload"

type ChapterService interface {
	Create(p payload.CreateChapter) (id string, err error)
}
