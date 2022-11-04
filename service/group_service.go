package service

import "github.com/erikrios/my-story-dummy/model/payload"

type GroupService interface {
	Create(p payload.CreateGroup) error
}
