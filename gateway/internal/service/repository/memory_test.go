package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bhongy/kimidori/gateway/internal/service/repository"
)

func TestMemory_ByServiceName(t *testing.T) {
	origin := "http://some-internal-addr:8000"

	m := repository.NewMemory(map[string]repository.Backend{
		"fake-svc": {Origin: origin},
	})

	be, err := m.ByServiceName("fake-svc")
	assert.NoError(t, err)
	assert.Equal(t, repository.Backend{Origin: origin}, be)
}
