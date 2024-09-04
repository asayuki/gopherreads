package library

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/asayuki/gopherreads/models"
)

func Scanner(rootpath string) ([]*models.Book, error) {
	books := make([]*models.Book, 0)

	err := filepath.Walk(rootpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip dotfiles
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// Check if directory or not
		if !info.IsDir() {
			// Not a directory, so check if its any known filetype
			for _, ext := range []string{".epub", ".zip", ".rar", ".cbz"} {
				if strings.EqualFold(filepath.Ext(info.Name()), ext) {
					name := info.Name()
					// fmt.Println(strings.TrimPrefix(path, rootpath))
					books = append(books, &models.Book{
						// Name:     strings.TrimSuffix(name, filepath.Ext(name)),
						Path:     strings.TrimPrefix(strings.TrimSuffix(path, name), rootpath),
						FullPath: strings.TrimPrefix(path, rootpath),
						Type:     strings.TrimPrefix(filepath.Ext(name), "."),
						Metadata: models.BookMeta{
							Title: strings.TrimSuffix(name, filepath.Ext(name)),
						},
					})
				}
			}
		} else {
			// It is a directory, so we check if all files inside it are images
			// to count it as a book/manga/etc
			onlyImages := true
			err = filepath.Walk(path, func(fp string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if path == fp {
					return nil
				}

				if !isImage(info.Name()) {
					onlyImages = false
					return filepath.SkipAll
				}

				return nil
			})

			if onlyImages {
				name := info.Name()
				// (strings.TrimPrefix(path, rootpath))
				books = append(books, &models.Book{
					// Name:     name,
					Path:     strings.TrimPrefix(strings.TrimSuffix(path, name), rootpath),
					FullPath: strings.TrimPrefix(path, rootpath),
					Type:     "folder",
					Metadata: models.BookMeta{
						Title: name,
					},
				})
			}

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return books, nil
}

func isImage(filename string) bool {
	switch filepath.Ext(filename) {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}
