package request

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validJSONBody   io.ReadCloser
	invalidJSONBody io.ReadCloser
	dest            interface{}
)

const (
	VALID_JSON_STRING   = "{\"success\": true}"
	INVALID_JSON_STRING = "{\"success\": ,,"
)

func TestParseJSON(t *testing.T) {
	initSuiteForParseJSON()

	testParseJSONFromBody(t, validJSONBody, true)
	testParseJSONFromBody(t, invalidJSONBody, false)
}

func initSuiteForParseJSON() {
	validJSONBody = ioutil.NopCloser(strings.NewReader(VALID_JSON_STRING))
	invalidJSONBody = ioutil.NopCloser(strings.NewReader(INVALID_JSON_STRING))
}

func testParseJSONFromBody(t *testing.T, body io.ReadCloser, validity bool) {
	err := ParseJSON(body, &dest)
	assert.Equal(t, validity, err == nil)
}
