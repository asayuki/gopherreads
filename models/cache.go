package models

import "time"

type Cache struct {
	Content    []byte
	Expiration time.Time
}

func (c *Cache) IsExpired() bool {
	return time.Now().After(c.Expiration)
}
