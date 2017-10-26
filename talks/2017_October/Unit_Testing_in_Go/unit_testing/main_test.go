// Author(s): Michael Koeppl

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	const expectedResult = 5
	// Use testify's assert (makes it simpler)
	assert.Equal(t, expectedResult, add(2, 3))
}
