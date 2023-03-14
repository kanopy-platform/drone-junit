package junit

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJunitFile(t *testing.T) {
	results, err := Read("testdata/report.xml")
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, results.Tests, 2)

	b, err := json.Marshal(results)
	assert.NoError(t, err)
	t.Log(string(b))

	want, err := os.ReadFile("testdata/sample.json")
	assert.NoError(t, err)
	tr := strings.TrimRight(string(want), "\n")
	assert.Equal(t, tr, string(b))
}
