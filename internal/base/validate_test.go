package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("validate true", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{UserID: "testUserID", Title: "testTitle", Description: "testDescription"}
		actual := base.Validate(&request)
		expect := make([]string, 0)
		assert.Equal(t, expect, actual)
	})

	t.Run("cursor nil validate false", func(t *testing.T) {
		t.Parallel()
		var cursor *base.ValidateRequest
		actual := base.Validate(cursor)
		expect := make([]string, 0)
		expect = append(expect, "ValidateRequest nil")
		assert.Equal(t, expect, actual)
	})

	t.Run("struct fields empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest UserID is empty",
			"ValidateRequest Title is empty",
			"ValidateRequest Description is empty"}
		assert.Equal(t, expect, actual)
	})

	t.Run("Description empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{UserID: "testUserID", Title: "testTitle"}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest Description is empty"}
		assert.Equal(t, expect, actual)
	})

	t.Run("Title and Description empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{UserID: "testUserID"}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest Title is empty",
			"ValidateRequest Description is empty"}
		assert.Equal(t, expect, actual)
	})

	t.Run("Title empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{UserID: "testUserID", Description: "testDescription"}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest Title is empty"}
		assert.Equal(t, expect, actual)
	})

	t.Run("UserID and Title empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{Description: "testDescription"}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest UserID is empty",
			"ValidateRequest Title is empty"}
		assert.Equal(t, expect, actual)
	})

	t.Run("UserID empty validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{Title: "testTitle", Description: "testDescription"}
		actual := base.Validate(&request)
		expect := []string{
			"ValidateRequest UserID is empty"}
		assert.Equal(t, expect, actual)
	})

}
