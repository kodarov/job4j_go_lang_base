package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
)

func TestTrackerGetItems(t *testing.T) {
	t.Parallel()

	t.Run("copy", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{ID: "1", Name: "First Item"}
		tr.AddItem(item)

		res := tr.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t, []tracker.Item{item}, tr.GetItems())
	})
}

func TestTrackerAddItem(t *testing.T) {
	t.Parallel()

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{ID: "1", Name: "First Item"}

		tr.AddItem(item)

		assert.Equal(t, []tracker.Item{item}, tr.GetItems())
	})

	t.Run("many items", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		first := tracker.Item{ID: "1", Name: "First Item"}
		second := tracker.Item{ID: "2", Name: "Second Item"}

		tr.AddItem(first)
		tr.AddItem(second)

		assert.Equal(t, []tracker.Item{first, second}, tr.GetItems())
	})
}
