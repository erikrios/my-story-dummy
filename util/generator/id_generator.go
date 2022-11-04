package generator

import (
	"fmt"
	"math/rand"
	"time"
)

type IDGenerator interface {
	GenerateStoryID() (id string)
	GenerateChapterID() (id string)
}

type simpleIDGenerator struct{}

func NewSimpleIDGenerator() *simpleIDGenerator {
	return &simpleIDGenerator{}
}

func (s *simpleIDGenerator) GenerateStoryID() (id string) {
	id = s.generate(4)
	id = fmt.Sprintf("s-%s", id)
	return
}

func (s *simpleIDGenerator) GenerateChapterID() (id string) {
	id = s.generate(6)
	id = fmt.Sprintf("c-%s", id)
	return
}

func (s *simpleIDGenerator) generate(n uint8) string {
	rand.Seed(time.Now().UnixNano())

	res := make([]byte, n)

	for i := 0; uint8(i) < n; i++ {
		x := rand.Intn(26)
		res[i] = 'a' + uint8(x)
	}

	return string(res)
}
