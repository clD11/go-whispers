package shared

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/clD11/go-whispers/golang/testsupport"

	"github.com/stretchr/testify/assert"
)

func TestWriteErrorResponse(t *testing.T) {
	rw := httptest.NewRecorder()
	code := testsupport.RandomInt()
	errors := []string{testsupport.RandomString(), testsupport.RandomString()}

	WriteErrorResponse(rw, code, errors[0], errors[1])

	var actual WrappedResponse
	_ = json.NewDecoder(rw.Body).Decode(&actual)

	assert.Equal(t, code, rw.Code)
	assert.Equal(t, errors, actual.Error)
}
