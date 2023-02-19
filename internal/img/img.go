package img

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

type Img struct {
	path  string
	files *[]os.DirEntry
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

func (i Img) GetRandomImagePath() string {
	path := i.path + "/" + (*i.files)[rand.Intn(len(*i.files))].Name()
	return path
}

func (i Img) GetRandomDog() (string, error) {
	get, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return "", err
	}
	if get.StatusCode == http.StatusOK {
		response := struct {
			Message string `json:"Message"`
			Status  string `json:"Status"`
		}{}
		all, err := io.ReadAll(get.Body)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(all, &response)
		if err != nil {
			return "", err
		}
		if response.Status != "success" || response.Message == "" {
			return "", fmt.Errorf("not success Status of api: '%s', Message: %s", response.Status, response.Message)
		}
		return response.Message, nil
	} else {
		return "", fmt.Errorf("not ok Status code: %d", get.StatusCode)
	}
}
