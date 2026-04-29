package base_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
	"testing"
)

func Test_Add(t *testing.T) {

	rsl := base.Add(1, 2)
	expected := 3

	assert.Equal(t, rsl, expected)

}
