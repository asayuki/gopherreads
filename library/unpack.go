package library

import (
	"fmt"

	"github.com/asayuki/gopherreads/models"
)

func Unpack(item models.Book, cache map[string]*models.Cache) ([]string, error) {
	switch item.Type {
	case "zip", "cbz":
		return UnpackArchive(item.Path, cache)
	case "epub":
		return UnpackEpub(item.Path, cache)
	case "folder":
		return UnpackFolder(item.Path, cache)
	default:
		return nil, fmt.Errorf("not a valid type to unpack")
	}
}

func UnpackArchive(path string, cache map[string]*models.Cache) ([]string, error) {
	return nil, nil
}

func UnpackEpub(path string, cache map[string]*models.Cache) ([]string, error) {
	return nil, nil
}

func UnpackFolder(path string, cache map[string]*models.Cache) ([]string, error) {
	return nil, nil
}
