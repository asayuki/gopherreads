package models

import "strings"

type Directory struct {
	Name        string
	Path        string
	Directories []*Directory
}

type DirectoryItem struct {
	Name string
	Path string
}

func (d *Directory) GetDirectories(path string) []DirectoryItem {
	current := d
	directories := make([]DirectoryItem, 0)

	path = strings.TrimPrefix(path, "/")
	if path != "" {
		parts := strings.Split(path, "/")
		for _, part := range parts {
			for _, dir := range current.Directories {
				if strings.EqualFold(dir.Name, part) {
					current = dir
					break
				}
			}

			if !strings.EqualFold(current.Name, parts[len(parts)-1]) {
				return nil
			}
		}
	}

	for _, p := range current.Directories {
		directories = append(directories, DirectoryItem{
			Name: p.Name,
			Path: p.Path,
		})
	}

	return directories
}
