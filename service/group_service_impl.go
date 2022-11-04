package service

import (
	"encoding/json"
	"fmt"

	"github.com/erikrios/my-story-dummy/entity"
	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/util/customerr"
	cfs "github.com/erikrios/my-story-dummy/util/fs"
	"github.com/erikrios/my-story-dummy/validation"
)

type groupServiceImpl struct {
	fs cfs.FS
}

func NewGroupServiceImpl(fs cfs.FS) *groupServiceImpl {
	return &groupServiceImpl{fs: fs}
}

func (g *groupServiceImpl) Create(p payload.CreateGroup) (err error) {
	if validateErr := validation.Validate.Struct(p); validateErr != nil {
		err = customerr.InvalidPayload
		return
	}

	absolutePath := fmt.Sprintf("%s/%s.json", "data", p.Name)

	if g.fs.IsExists(absolutePath) {
		err = customerr.ResourceAlreadyExists
		return
	}

	if dirErr := g.fs.CreateFile(absolutePath); dirErr != nil {
		err = customerr.Internal
		return
	}

	resp := entity.Response{
		Status:  "success",
		Message: "successfully get the stories",
		Data:    []entity.Story{},
	}

	respByte, marshalErr := json.MarshalIndent(resp, "", "  ")
	if marshalErr != nil {
		err = customerr.Internal
		return
	}

	if writeErr := g.fs.WriteFile(absolutePath, respByte); writeErr != nil {
		err = customerr.Internal
		return
	}

	return
}
