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

	// stores Background vars
	Pos *stack.Stack

	cursor int
}

func getDirs(dir string) ([]string, error) {
	_dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directories: %v\n", err)
	}

	var dirs []string
	var name string
	for _, d := range _dirs {
		name = d.Name()
		if d.IsDir() {
			name += "/"
		}
		dirs = append(dirs, name)
	}

	return dirs, nil
}

func (b *Background) MoveTo(dirname string) {
	var err error
	/**
	 * push current Background to Pos
	 * set current cursor to  zero
	 * set current directory list based on passed dirname
	 */
	b.Pos.Push(*b) // this also stores the cursor which is just awesome
	b.cursor = 0
	b.CurrDirs, err = getDirs(dirname)
	b.CurrDirs = append([]string{".", ".."}, b.CurrDirs...)
	if err != nil {
	}
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
		Cwd:         d,
		ForwardDirs: nil,
		CurrDirs:    append([]string{".", ".."}, currdirs...),
		Pos:         &stack.Stack{},
		cursor:      0,
	}, nil
}
