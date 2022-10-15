//go:build unit
// +build unit

package metrics

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/sudoblockio/icon-transformer/config"

	"github.com/stretchr/testify/assert"
)

func TestMetricsApiStart(t *testing.T) {
	assert := assert.New(t)

	// Set env
	os.Setenv("METRICS_PORT", "8888")
	os.Setenv("METRICS_PREFIX", "/metrics")

	config.ReadEnvironment()

	// Start metrics server
	Start()
	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s%s", config.Config.MetricsPort, config.Config.MetricsPrefix))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}

func TestCreateMetric(t *testing.T) {
	metrics := CreateGuage("foo", "bar", map[string]string{"type": "foo", "this": "that"})
	assert.Equal(t, metrics.Desc().String()[:4], "Desc")
	// TODO: How to get at `(*metrics.Desc()).help`
}
