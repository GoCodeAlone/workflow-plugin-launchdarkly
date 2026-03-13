package internal

import (
	"context"
	"testing"
)

func TestModuleInit_RegistersClient(t *testing.T) {
	m, err := newLaunchDarklyModule("test-init", map[string]any{
		"apiKey": "test-api-key",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-init")
	if !ok || c == nil {
		t.Error("expected client to be registered")
	}
	// cleanup
	UnregisterClient("test-init")
}

func TestModuleStop_UnregistersClient(t *testing.T) {
	m, _ := newLaunchDarklyModule("test-stop", map[string]any{
		"apiKey": "test-api-key",
	})
	_ = m.Init()
	_ = m.Stop(context.Background())
	_, ok := GetClient("test-stop")
	if ok {
		t.Error("expected client to be unregistered after stop")
	}
}

func TestModuleInit_MissingAPIKey(t *testing.T) {
	m, err := newLaunchDarklyModule("test-missing", map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err == nil {
		t.Error("expected error for missing apiKey")
		UnregisterClient("test-missing")
	}
}

func TestModuleInit_WithCustomURL(t *testing.T) {
	m, err := newLaunchDarklyModule("test-custom-url", map[string]any{
		"apiKey": "test-api-key",
		"apiUrl": "https://custom.example.com/api/v2",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-custom-url")
	if !ok || c == nil {
		t.Error("expected client to be registered with custom URL")
	}
	if c.baseURL != "https://custom.example.com/api/v2" {
		t.Errorf("expected custom URL, got %q", c.baseURL)
	}
	UnregisterClient("test-custom-url")
}

func TestModuleInit_DefaultBaseURL(t *testing.T) {
	m, err := newLaunchDarklyModule("test-default-url", map[string]any{
		"apiKey": "test-api-key",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-default-url")
	if !ok || c == nil {
		t.Fatal("expected client to be registered")
	}
	if c.baseURL != defaultBaseURL {
		t.Errorf("expected default URL %q, got %q", defaultBaseURL, c.baseURL)
	}
	UnregisterClient("test-default-url")
}
