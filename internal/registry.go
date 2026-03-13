package internal

import "sync"

var (
	clientMu       sync.RWMutex
	clientRegistry = make(map[string]*launchDarklyClient)
)

// RegisterClient adds a LaunchDarkly client to the global registry under the given name.
func RegisterClient(name string, c *launchDarklyClient) {
	clientMu.Lock()
	defer clientMu.Unlock()
	clientRegistry[name] = c
}

// GetClient looks up a LaunchDarkly client by name.
func GetClient(name string) (*launchDarklyClient, bool) {
	clientMu.RLock()
	defer clientMu.RUnlock()
	c, ok := clientRegistry[name]
	return c, ok
}

// UnregisterClient removes a client from the registry.
func UnregisterClient(name string) {
	clientMu.Lock()
	defer clientMu.Unlock()
	delete(clientRegistry, name)
}
