package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// relayConfigListStep implements step.launchdarkly_relay_config_list
type relayConfigListStep struct {
	name       string
	moduleName string
}

func newRelayConfigListStep(name string, config map[string]any) (*relayConfigListStep, error) {
	return &relayConfigListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *relayConfigListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/relay-proxy-configs", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// relayConfigGetStep implements step.launchdarkly_relay_config_get
type relayConfigGetStep struct {
	name       string
	moduleName string
}

func newRelayConfigGetStep(name string, config map[string]any) (*relayConfigGetStep, error) {
	return &relayConfigGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *relayConfigGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	relayID := resolveValue("relayId", current, config)
	if relayID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "relayId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/relay-proxy-configs/%s", relayID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// relayConfigCreateStep implements step.launchdarkly_relay_config_create
type relayConfigCreateStep struct {
	name       string
	moduleName string
}

func newRelayConfigCreateStep(name string, config map[string]any) (*relayConfigCreateStep, error) {
	return &relayConfigCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *relayConfigCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	relayName := resolveValue("name", current, config)
	if relayName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	body := map[string]any{"name": relayName}
	if policy := resolveMap("policy", current, config); policy != nil {
		body["policy"] = policy
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/relay-proxy-configs", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// relayConfigUpdateStep implements step.launchdarkly_relay_config_update
type relayConfigUpdateStep struct {
	name       string
	moduleName string
}

func newRelayConfigUpdateStep(name string, config map[string]any) (*relayConfigUpdateStep, error) {
	return &relayConfigUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *relayConfigUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	relayID := resolveValue("relayId", current, config)
	if relayID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "relayId is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/relay-proxy-configs/%s", relayID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// relayConfigDeleteStep implements step.launchdarkly_relay_config_delete
type relayConfigDeleteStep struct {
	name       string
	moduleName string
}

func newRelayConfigDeleteStep(name string, config map[string]any) (*relayConfigDeleteStep, error) {
	return &relayConfigDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *relayConfigDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	relayID := resolveValue("relayId", current, config)
	if relayID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "relayId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/relay-proxy-configs/%s", relayID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
