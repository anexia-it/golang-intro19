package intro19

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTomlFormat_Decode(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		tomlConverter := tomlFormat{}
		input := `name = "test"
version = "1.0.0"

[object]
  key1 = "value 1"
  key2 = "value 2"
`

		expected := map[string]interface{}{
			"name":    "test",
			"version": "1.0.0",
			"object": map[string]interface{}{
				"key1": "value 1",
				"key2": "value 2",
			},
		}
		output := make(map[string]interface{})
		err := tomlConverter.Decode(strings.NewReader(input), &output)
		assert.NoError(t, err)
		assert.EqualValues(t, expected, output)

	})

	t.Run("Error", func(t *testing.T) {
		tomlConverter := tomlFormat{}

		input := "---"

		output := make(map[string]interface{})
		err := tomlConverter.Decode(strings.NewReader(input), &output)
		assert.Error(t, err)
	})
}

func TestTomlFormat_Encode(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		in := map[string]interface{}{
			"key1": "value 1",
			"object": map[string]interface{}{
				"key2": "value 2",
			},
		}

		expected := `key1 = "value 1"

[object]
  key2 = "value 2"
`
		tomlConverter := tomlFormat{}
		rw := httptest.NewRecorder()

		err := tomlConverter.Encode(rw, in)
		assert.NoError(t, err)
		assert.EqualValues(t, "application/toml", rw.HeaderMap.Get("Content-Type"))
		assert.EqualValues(t, expected, rw.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		tomlConverter := tomlFormat{}
		in := map[int]string{
			1: "1",
		}
		rw := httptest.NewRecorder()

		err := tomlConverter.Encode(rw, in)
		assert.Error(t, err)
	})
}
