package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {
	p := Name{}
	assert.Equal(t, p.String(), "")
	p = p.GetChild("http")
	assert.Equal(t, p.String(), "http")
	p = p.GetChild("request")
	assert.Equal(t, p.String(), "http.request")
	q := ParseName("http.request")
	assert.Equal(t, p, q)
}
