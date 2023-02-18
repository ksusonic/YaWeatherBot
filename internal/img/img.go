package img

import (
	"fmt"
	"math/rand"
	"os"
)

type Img struct {
	path  string
	files *[]os.DirEntry
}

func (i Img) GetRandomImagePath() string {
	path := i.path + "/" + (*i.files)[rand.Intn(len(*i.files))].Name()
	return path
}

func NewImg(path string) (*Img, error) {
	if path == "" {
		return nil, fmt.Errorf("empty img path")
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("could not find dir: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("0 files in dir")
	}
	return &Img{
		path:  path,
		files: &files,
	}, nil
}
