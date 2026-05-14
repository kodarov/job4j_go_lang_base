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
		_, err := tr.AddItem(item)
		assert.NoError(t, err)

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

		_, err := tr.AddItem(item)
		assert.NoError(t, err)

		assert.Equal(t, []tracker.Item{item}, tr.GetItems())
	})

	t.Run("many items", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		first := tracker.Item{ID: "1", Name: "First Item"}
		second := tracker.Item{ID: "2", Name: "Second Item"}

		_, err := tr.AddItem(first)
		assert.NoError(t, err)

		_, err = tr.AddItem(second)
		assert.NoError(t, err)

		assert.Equal(t, []tracker.Item{first, second}, tr.GetItems())
	})

	t.Run("error add - already exists", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		first := tracker.Item{ID: "1", Name: "First Item"}
		second := tracker.Item{ID: "1", Name: "Second Item"}

		_, err := tr.AddItem(first)
		assert.NoError(t, err)

		_, err = tr.AddItem(second)
		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
		assert.Equal(t, []tracker.Item{first}, tr.GetItems())
	})
}

func TestTrackerUpdateItem(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tr.UpdateItem(item)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})

}
