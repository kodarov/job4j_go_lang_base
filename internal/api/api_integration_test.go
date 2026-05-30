package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"job4j.ru/go-lang-base/internal/api"
	"job4j.ru/go-lang-base/internal/repository"
	"job4j.ru/go-lang-base/internal/testhelpers"
	"job4j.ru/go-lang-base/internal/tracker"
)

var app *fiber.App
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
	server := api.NewServer(repo)

	app = fiber.New()
	server.Route(app.Group("/api"))

	os.Exit(m.Run())
}

func setupTest(t *testing.T) {
	err := testhelpers.CleanDatabase(ctx, repo.Pool())
	require.NoError(t, err)
}

func TestAPI_CreateItem(t *testing.T) {
	setupTest(t)

	item := tracker.Item{
		ID:   uuid.New().String(),
		Name: "API Item",
	}
	body, _ := json.Marshal(item)

	req := httptest.NewRequest(http.MethodPost, "/api/items/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var created tracker.Item
	err = json.NewDecoder(resp.Body).Decode(&created)
	assert.NoError(t, err)
	assert.Equal(t, item.Name, created.Name)
	assert.NotEmpty(t, created.ID)
}

func TestAPI_GetItem(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "To Get"}
	require.NoError(t, repo.Create(ctx, item))

	req := httptest.NewRequest(http.MethodGet, "/api/items/"+item.ID, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var found tracker.Item
	err = json.NewDecoder(resp.Body).Decode(&found)
	assert.NoError(t, err)
	assert.Equal(t, item, found)
}

func TestAPI_UpdateItem(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "Original"}
	require.NoError(t, repo.Create(ctx, item))

	update := tracker.Item{Name: "Updated"}
	body, _ := json.Marshal(update)

	req := httptest.NewRequest(http.MethodPatch, "/api/items/"+item.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated tracker.Item
	err = json.NewDecoder(resp.Body).Decode(&updated)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updated.Name)
}

func TestAPI_DeleteItem(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "To Delete"}
	require.NoError(t, repo.Create(ctx, item))

	req := httptest.NewRequest(http.MethodDelete, "/api/items/"+item.ID, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	_, err = repo.Get(ctx, item.ID)
	assert.ErrorIs(t, err, tracker.ErrNotFound)
}

func TestAPI_GetItem_NotFound(t *testing.T) {
	setupTest(t)

	req := httptest.NewRequest(http.MethodGet, "/api/items/"+uuid.New().String(), nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAPI_GetItems(t *testing.T) {
	setupTest(t)

	item1 := tracker.Item{ID: uuid.New().String(), Name: "Item 1"}
	item2 := tracker.Item{ID: uuid.New().String(), Name: "Item 2"}
	require.NoError(t, repo.Create(ctx, item1))
	require.NoError(t, repo.Create(ctx, item2))

	req := httptest.NewRequest(http.MethodGet, "/api/items/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result struct {
		Items []tracker.Item `json:"items"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Len(t, result.Items, 2)
	assert.Contains(t, result.Items, item1)
	assert.Contains(t, result.Items, item2)
}

func TestAPI_CreateItem_InvalidJSON(t *testing.T) {
	setupTest(t)

	req := httptest.NewRequest(http.MethodPost, "/api/items/", bytes.NewReader([]byte(`{"name": "missing quote}`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAPI_UpdateItem_InvalidJSON(t *testing.T) {
	setupTest(t)

	item := tracker.Item{ID: uuid.New().String(), Name: "To Update"}
	require.NoError(t, repo.Create(ctx, item))

	req := httptest.NewRequest(http.MethodPatch, "/api/items/"+item.ID, bytes.NewReader([]byte(`{"name": "missing quote}`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
