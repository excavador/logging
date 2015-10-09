package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	http_request := GetLogger("http.request")
	assert.Equal(t, NewName("http.request"), http_request.Name())
	db := GetLogger("db")
	assert.Equal(t, NewName("db"), db.Name())
}
