package stores

import (
	"database/sql"
	"fmt"

	"github.com/asayuki/gopherreads/models"
)

type LibraryStore struct {
	db *sql.DB
}

func NewLibraryStore(db *sql.DB) *LibraryStore {
	return &LibraryStore{db}
}

func (s *LibraryStore) GetBookByPath(path string) (*models.Book, error) {
	rows := s.db.QueryRow(`
		SELECT
			b.id,
			b.path,
			b.full_path,
			b.name,
			b.type,
			b.scanned_at,
			b.last_scanned_at,
			COALESCE(bm.title, '') as title,
			COALESCE(bm.description, '') as description,
			COALESCE(bm.cover, '') as cover,
			COALESCE(bm.genre, '') as genre,
			COALESCE(bm.author, '') as author
		FROM books b
		LEFT JOIN bookmeta bm ON b.id = bm.book_id
		WHERE b.full_path = ?
	`, path)

	book := new(models.Book)
	err := scanBook(rows, book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *LibraryStore) GetBooks() ([]*models.Book, error) {
	rows, err := s.db.Query(`
		SELECT
			b.id,
			b.path,
			b.full_path,
			b.name,
			b.type,
			b.scanned_at,
			b.last_scanned_at,
			COALESCE(bm.title, '') as title,
			COALESCE(bm.description, '') as description,
			COALESCE(bm.cover, '') as cover,
			COALESCE(bm.genre, '') as genre,
			COALESCE(bm.author, '') as author
		FROM books b
		LEFT JOIN bookmeta bm on b.id = bm.book_id
	`)

	if err != nil {
		return nil, err
	}

	books, err := scanBooks(rows)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *LibraryStore) GetBooksByPath(path string) ([]*models.Book, error) {
	rows, err := s.db.Query(`
		SELECT
			b.id,
			b.path,
			b.full_path,
			b.name,
			b.type,
			b.scanned_at,
			b.last_scanned_at,
			COALESCE(bm.title, '') as title,
			COALESCE(bm.description, '') as description,
			COALESCE(bm.cover, '') as cover,
			COALESCE(bm.genre, '') as genre,
			COALESCE(bm.author, '') as author
		FROM books b
		LEFT JOIN bookmeta bm on b.id = bm.book_id
		WHERE b.path = ?
	`, path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	books, err := scanBooks(rows)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return books, nil
}

func (s *LibraryStore) InsertBook(book models.Book) (int64, error) {
	result, err := s.db.Exec(`
		INSERT INTO books (
			path, full_path, name, type
		) VALUES (?, ?, ?, ?)
	`, book.Path, book.FullPath, book.Name, book.Type)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LibraryStore) InsertBookMeta(id int, bookmeta models.BookMeta) error {
	_, err := s.db.Exec(`
		INSERT INTO bookmeta (
			book_id, title, description, cover, genre, author
		) VALUES(?, ?, ?, ?, ?, ?)
	`, id, bookmeta.Title, bookmeta.Description, bookmeta.Cover, bookmeta.Cover, bookmeta.Genre, bookmeta.Author)

	return err
}

func (s *LibraryStore) UpdateBookMeta(id, bookmeta models.BookMeta) error {
	_, err := s.db.Exec(`
		UPDATE bookmeta (
			title, description, cover, genre, author
		), VALUES(?, ?, ?, ?, ?)
		WHERE book_id = ?
	`, bookmeta.Title, bookmeta.Description, bookmeta.Cover, bookmeta.Cover, bookmeta.Genre, bookmeta.Author, id)

	return err
}

func scanBook(row *sql.Row, book *models.Book) error {
	err := row.Scan(
		&book.ID,
		&book.Path,
		&book.FullPath,
		&book.Name,
		&book.Type,
		&book.ScannedAt,
		&book.LastScannedAt,
		&book.Metadata.Author,
		&book.Metadata.Cover,
		&book.Metadata.Description,
		&book.Metadata.Genre,
		&book.Metadata.Title,
	)

	return err
}

func scanBooks(rows *sql.Rows) ([]*models.Book, error) {
	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Path,
			&book.FullPath,
			&book.Name,
			&book.Type,
			&book.ScannedAt,
			&book.LastScannedAt,
			&book.Metadata.Author,
			&book.Metadata.Cover,
			&book.Metadata.Description,
			&book.Metadata.Genre,
			&book.Metadata.Title,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
