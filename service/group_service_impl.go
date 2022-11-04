package service

import (
	"fmt"

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
		err = customerr.PayloadParse
		return
	}

	absolutePath := fmt.Sprintf("%s/%s", "data", p.Name)

	if g.fs.IsExists(absolutePath) {
		err = customerr.ResourceAlreadyExists
		return
	}

	if dirErr := g.fs.CreateDir(absolutePath); dirErr != nil {
		err = customerr.Internal
		return
	}

	return
}
