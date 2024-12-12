package testdata

import (
	"github.com/guregu/null/v5"
)

type b struct {
	a string
}

type Model1 struct {
	b
	name     string         `gs:"rw"`
	age      int            `gs:"rw"`
	nullable null.Bool      `gs:"rw"`
	readonly []byte         `gs:"r"`
	pointer  *byte          `gs:"w"`
	m        map[string]int `gs:"r"`
	c        chan struct{}  `gs:"w"`
	s        b              `gs:"w"`
}

func (m *Model1) Create(a []string) {

}
