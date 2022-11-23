package __test__

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	client := newClient()
	res, err := client.getUsers()
	if err != nil {
		t.Fatalf("failed because: %s\n", err.Error())
	}

	assert.Equal(t, 200, res.StatusCode)
}
