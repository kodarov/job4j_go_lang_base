package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestLruCachePut(t *testing.T) {
	t.Parallel()

	t.Run("full", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("hello1", "world1")
		cache.Put("hello2", "world2")
		cache.Put("hello3", "world3")
		cache.Put("hello4", "world4")

		assert.Equal(t, (*string)(nil), cache.Get("hello1"))
		assert.Equal(t, "world4", *cache.Get("hello4"))
	})

	t.Run("size one", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)

		cache.Put("hello1", "world1")
		cache.Put("hello2", "world2")
		cache.Put("hello3", "world3")
		cache.Put("hello4", "world4")

		assert.Nil(t, cache.Get("hello1"))
		assert.Nil(t, cache.Get("hello2"))
		assert.Nil(t, cache.Get("hello3"))
		assert.Equal(t, "world4", *cache.Get("hello4"))
	})

	t.Run("zero size", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(0)

		cache.Put("hello1", "world1")

		assert.Nil(t, cache.Get("hello1"))
	})
}

func TestLruCacheGet(t *testing.T) {
	t.Parallel()

	t.Run("one element", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(10)

		cache.Put("hello1", "world1")

		assert.Equal(t, "world1", *cache.Get("hello1"))
	})

	t.Run("missing", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("hello1", "world1")

		assert.Nil(t, cache.Get("unknown"))
	})

	t.Run("head", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("hello1", "world1")
		cache.Put("hello2", "world2")
		cache.Put("hello3", "world3")
		assert.Equal(t, "world3", *cache.Get("hello3"))
		cache.Put("hello4", "world4")

		assert.Nil(t, cache.Get("hello1"))
		assert.Equal(t, "world2", *cache.Get("hello2"))
		assert.Equal(t, "world3", *cache.Get("hello3"))
		assert.Equal(t, "world4", *cache.Get("hello4"))
	})

	t.Run("middle", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("hello1", "world1")
		cache.Put("hello2", "world2")
		cache.Put("hello3", "world3")
		assert.Equal(t, "world2", *cache.Get("hello2"))
		cache.Put("hello4", "world4")

		assert.Nil(t, cache.Get("hello1"))
		assert.Equal(t, "world2", *cache.Get("hello2"))
		assert.Equal(t, "world3", *cache.Get("hello3"))
		assert.Equal(t, "world4", *cache.Get("hello4"))
	})

	t.Run("tail", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("hello1", "world1")
		cache.Put("hello2", "world2")
		cache.Put("hello3", "world3")
		assert.Equal(t, "world1", *cache.Get("hello1"))
		cache.Put("hello4", "world4")

		assert.Equal(t, "world1", *cache.Get("hello1"))
		assert.Nil(t, cache.Get("hello2"))
		assert.Equal(t, "world3", *cache.Get("hello3"))
		assert.Equal(t, "world4", *cache.Get("hello4"))
	})
}
