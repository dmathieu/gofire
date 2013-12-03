package gofire

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulNew(t *testing.T) {
	client := NewClient("1234", "gofire")
	assert.Equal(t, "1234", client.token)
	assert.Equal(t, "https://gofire.campfirenow.com", client.baseURL)
}
