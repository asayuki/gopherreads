package library

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
)

func Categorize(librarystore stores.LibraryStore, directories *models.Directory) {
	books, err := Scanner(config.Envs.LibraryPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, book := range books {
		// fmt.Println(book.Path)
		// fmt.Println(book.FullPath)
		// fmt.Println(book.Type)
		currentDirectory := directories
		for _, part := range strings.Split(book.Path, "/") {
			found := false
			for _, dir := range currentDirectory.Directories {
				if dir.Name == part {
					currentDirectory = dir
					found = true
					break
				}
			}

			if !found {
				newDirectory := &models.Directory{
					Name: part,
					Path: filepath.Join(currentDirectory.Path, part),
				}
				currentDirectory.Directories = append(currentDirectory.Directories, newDirectory)
				currentDirectory = newDirectory
			}
		}

		_, err := librarystore.GetBookByPath(book.FullPath)
		if err != nil {
			fmt.Println(err)
			if err == sql.ErrNoRows {
				librarystore.InsertBook(*book)
			}
		}
	}

	// fmt.Println(directories.GetDirectories("asdasd"))
}
