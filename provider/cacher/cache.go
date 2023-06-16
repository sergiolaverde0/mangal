package cacher

import (
	"github.com/belphemur/mangal/filesystem"
	"github.com/belphemur/mangal/util"
	"github.com/belphemur/mangal/where"
	"github.com/metafates/gache"
	"github.com/samber/mo"
	"path/filepath"
	"time"
)

type Cacher[T any] struct {
	internal *gache.Cache[map[string]T]
}

func NewCacher[T any](name string, ttl time.Duration) *Cacher[T] {
	return &Cacher[T]{
		internal: gache.New[map[string]T](&gache.Options{
			Lifetime:   ttl,
			Path:       filepath.Join(where.Cache(), util.SanitizeFilename(name)+".json"),
			FileSystem: &filesystem.GacheFs{},
		}),
	}
}

func (c *Cacher[T]) Get(key string) mo.Option[T] {
	data, expired, err := c.internal.Get()
	if err != nil || expired || data == nil {
		return mo.None[T]()
	}

	if x, ok := data[key]; ok {
		return mo.Some(x)
	}

	return mo.None[T]()
}

func (c *Cacher[T]) Set(key string, t T) error {
	data, expired, err := c.internal.Get()

	if err != nil {
		return err
	}

	if expired || data == nil {
		data = make(map[string]T)
	}

	data[key] = t

	return c.internal.Set(data)
}
