package library

import (
	"database/sql"
	"fmt"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
)

func Categorize(librarystore stores.LibraryStore, directories *models.Directory) {
	books, err := Scanner(config.Envs.LibraryPath)
	if err != nil {
		fmt.Println(err)
	}

	directories.Mu.Lock()
	directories.Directories = nil
	directories.Mu.Unlock()

	for _, book := range books {
		directories.AddDirectory(book.Path)

		_, err := librarystore.GetBookByPath(book.FullPath)
		if err != nil {
			if err == sql.ErrNoRows {
				librarystore.InsertBook(*book)
			}
		}
	}

	// Todo: remove old books.
}
