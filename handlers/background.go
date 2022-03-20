package handlers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang-collections/collections/stack"
)

type Background struct {
	Cwd         string
	CurrDirs    []string
	ForwardDirs []string
	BackDirs    []string

	Pos *stack.Stack

	cursor        int
	selection int
}

func getDirs(dir string) ([]string, error) {
	_dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directories: %v\n", err)
	}

	var dirs []string
	for _, d := range _dirs {
		dirs = append(dirs, d.Name())
	}

	return dirs, nil
}

func (b *Background) MoveTo(dirname string) {
    var err error
    b.cursor = 0
    b.Cwd = dirname
    b.CurrDirs, err = getDirs(dirname)
    if err != nil {}
    b.Pos.Push(dirname)
}

func New() (*Background, error) {
	d, e := os.Getwd()
	if e != nil {
		return nil, fmt.Errorf("new model creation failed: %v\n", e)
	}

	currdirs, e := getDirs(d)
	if e != nil {
		return nil, fmt.Errorf("new model creation failed: %v\n", e)
	}

	return &Background{
		Cwd:           d,
		BackDirs:      nil,
		ForwardDirs:   nil,
		CurrDirs:      currdirs,
		Pos:           &stack.Stack{},
		cursor:        0,
		selection: -1,
	}, nil
}
