//go:build debugtest
// +build debugtest

package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.io/mondoo/motor/providers"
	gcp_transport "go.mondoo.io/mondoo/motor/providers/gcp"
)

func TestGcpDiscovery(t *testing.T) {
	projectid, err := gcp_transport.GetCurrentProject()
	require.NoError(t, err)

	pCfg := &providers.Config{
		Type: providers.ProviderType_GCP,
		Options: map[string]string{
			"project": projectid,
		},
		Discover: &providers.Discovery{
			Targets: []string{"all"},
		},
	}

	r := GcpProjectResolver{}
	assets, err := r.Resolve(pCfg)
	require.NoError(t, err)
	assert.True(t, len(assets) > 0)
}
