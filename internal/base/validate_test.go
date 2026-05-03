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
		expect = append(expect, "validate message")
		assert.Equal(t, expect, actual)
	})

	t.Run("cursor nil validate false", func(t *testing.T) {
		t.Parallel()
		var cursor *base.ValidateRequest
		actual := base.Validate(cursor)
		expect := make([]string, 0)
		expect = append(expect, "validate error")
		assert.Equal(t, expect, actual)
	})

	t.Run("struct nil validate false", func(t *testing.T) {
		t.Parallel()
		request := base.ValidateRequest{}
		actual := base.Validate(&request)
		expect := make([]string, 0)
		expect = append(expect, "validate error")
		assert.Equal(t, expect, actual)
	})
}
