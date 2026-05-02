package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Mono(t *testing.T) {
	t.Parallel()

	t.Run("[1, 2, 3] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[1, 1, 1] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 1, 1}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[3, 2, 1] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{3, 2, 1}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[3, 2, 4] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{3, 2, 4}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})

	t.Run("[1, 2, 1] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 1}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})
}
