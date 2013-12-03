package gofire

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulNew(t *testing.T) {
	client := NewClient("1234", "gofire", true)
	assert.Equal(t, "1234", client.token)
	assert.Equal(t, "https://gofire.campfirenow.com", client.baseURL)
}

func TestSuccessfulInsecureNew(t *testing.T) {
	client := NewClient("456", "gofire", false)
	assert.Equal(t, "456", client.token)
	assert.Equal(t, "http://gofire.campfirenow.com", client.baseURL)
}
