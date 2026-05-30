package repository_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"job4j.ru/go-lang-base/internal/repository"
	"job4j.ru/go-lang-base/internal/testhelpers"
	"job4j.ru/go-lang-base/internal/tracker"
)

var repo *repository.RepoPg
var ctx = context.Background()

func TestMain(m *testing.M) {
	pool, err := testhelpers.NewTestPool(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	err = testhelpers.Migrate(context.Background(), pool, "../../migrations")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not migrate database: %v\n", err)
		os.Exit(1)
	}

	repo = repository.NewRepoPg(pool)
	os.Exit(m.Run())
}

func setupTest(t *testing.T) {
	err := testhelpers.CleanDatabase(ctx, repo.Pool())
	require.NoError(t, err)
}

func TestRepoPg_CreateAndGet(t *testing.T) {
	setupTest(t)

	item := tracker.Item{
		ID:   uuid.New().String(),
		Name: "Test Item",
	}

	err := repo.Create(ctx, item)
	assert.NoError(t, err)

	found, err := repo.Get(ctx, item.ID)
	assert.NoError(t, err)
	assert.Equal(t, item, found)
}

func TestRepoPg_List(t *testing.T) {
	setupTest(t)

	item1 := tracker.Item{ID: uuid.New().String(), Name: "Item 1"}
	item2 := tracker.Item{ID: uuid.New().String(), Name: "Item 2"}

	require.NoError(t, repo.Create(ctx, item1))
	require.NoError(t, repo.Create(ctx, item2))

	items, err := repo.List(ctx)
	assert.NoError(t, err)
	assert.Contains(t, items, item1)
	assert.Contains(t, items, item2)
}

func TestRepoPg_Update(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "Original"}
	require.NoError(t, repo.Create(ctx, item))

	item.Name = "Updated"
	updated, err := repo.Update(ctx, item)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updated.Name)

	found, err := repo.Get(ctx, item.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", found.Name)
}

func TestRepoPg_Delete(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "To Delete"}
	require.NoError(t, repo.Create(ctx, item))

	err := repo.Delete(ctx, item.ID)
	assert.NoError(t, err)

	_, err = repo.Get(ctx, item.ID)
	assert.ErrorIs(t, err, tracker.ErrNotFound)
}
