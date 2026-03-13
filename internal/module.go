package internal

import (
	"context"
	"fmt"
)

// launchDarklyModule creates an HTTP client for the LaunchDarkly API and registers it.
type launchDarklyModule struct {
	name   string
	config map[string]any
}

func newLaunchDarklyModule(name string, config map[string]any) (*launchDarklyModule, error) {
	return &launchDarklyModule{name: name, config: config}, nil
}

// Init creates the LaunchDarkly HTTP client and registers it in the global registry.
func (m *launchDarklyModule) Init() error {
	apiKey, _ := m.config["apiKey"].(string)
	if apiKey == "" {
		return fmt.Errorf("launchdarkly.provider %q: apiKey is required", m.name)
	}
	apiURL, _ := m.config["apiUrl"].(string)

	client := newLaunchDarklyHTTPClient(apiKey, apiURL)
	RegisterClient(m.name, client)
	return nil
}

// Start is a no-op for this module.
func (m *launchDarklyModule) Start(_ context.Context) error { return nil }

// Stop unregisters the LaunchDarkly client.
func (m *launchDarklyModule) Stop(_ context.Context) error {
	UnregisterClient(m.name)
	return nil
}
