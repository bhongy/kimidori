package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileSystem implements Repository interface using local file system
type FileSystem struct {
	dir string
}

// NewFileSystem creates a new file-system repository
//
// `dir` is the path where the .json files for services can be found
// which can be relative (to the binary cwd) or absolute
func NewFileSystem(dir string) FileSystem {
	return FileSystem{dir: dir}
}

func (fs FileSystem) ByServiceName(svc string) (Backend, error) {
	p := filepath.Join(fs.dir, svc+".json")
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return Backend{}, ErrNotFound
	}

	b, err := ioutil.ReadFile(p)
	if err != nil {
		err = fmt.Errorf("cannot read service file: service=%s, path=%s, error: %v", svc, p, err)
		return Backend{}, err
	}

	var be Backend
	if err = json.Unmarshal(b, &be); err != nil {
		err = fmt.Errorf("cannot unmarshal service data from JSON file: path=%s, error: %v", p, err)
		return Backend{}, err
	}

	return be, nil
}
