package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {
	p := path{}
	assert.Equal(t, p.String(), "")
	p = p.GetChild("http")
	assert.Equal(t, p.String(), "http")
	p = p.GetChild("request")
	assert.Equal(t, p.String(), "http.request")
	q := parsePath("http.request")
	assert.Equal(t, p, q)
}
