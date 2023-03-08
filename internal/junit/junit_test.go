package junit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJunitFile(t *testing.T) {
	results, err := Read("testdata/report.xml")
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, results.Tests, 2)
}
