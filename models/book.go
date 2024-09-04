package models

import "time"

type Book struct {
	ID            int       `json:""`
	Path          string    `json:"path"`
	FullPath      string    `json:"full_path"`
	Type          string    `json:"type"`
	Metadata      BookMeta  `json:"metadata"`
	ScannedAt     time.Time `json:"scanned_at"`
	LastScannedAt time.Time `json:"last_scanned_at"`
}

type BookMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Genre       string `json:"genre"`
	Author      string `json:"author"`
}
