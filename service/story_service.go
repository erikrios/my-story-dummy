package service

import "github.com/erikrios/my-story-dummy/model/payload"

type StoryService interface {
	Create(p payload.CreateStory) (id string, err error)
}
