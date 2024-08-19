package models

import (
	"path/filepath"
	"strings"
	"sync"
)

type Directory struct {
	Name        string
	Path        string
	Directories []*Directory
	Mu          sync.RWMutex
}

type DirectoryItem struct {
	Name string
	Path string
}

func (d *Directory) AddDirectory(path string) {
	d.Mu.Lock()
	defer d.Mu.Unlock()

	parts := strings.Split(filepath.Clean(path), string(filepath.Separator))
	current := d

	for _, part := range parts {
		found := false
		for _, dir := range current.Directories {
			if strings.EqualFold(dir.Name, part) {
				current = dir
				found = true
				break
			}
		}
		if !found {
			newDir := &Directory{
				Name: part,
				Path: filepath.Join(current.Path, part),
			}
			current.Directories = append(current.Directories, newDir)
			current = newDir
		}
	}
}

func (d *Directory) GetDirectories(path string) []DirectoryItem {
	d.Mu.Lock()
	defer d.Mu.Unlock()

	parts := strings.Split(filepath.Clean(path), string(filepath.Separator))
	current := d

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}

		found := false

		for _, dir := range current.Directories {
			if strings.EqualFold(dir.Name, part) {
				current = dir
				found = true
				break
			}
		}

		if !found {
			return nil
		}
	}

	var directories []DirectoryItem
	for _, dir := range current.Directories {
		directories = append(directories, DirectoryItem{
			Name: dir.Name,
			Path: dir.Path,
		})
	}

	return directories
}
