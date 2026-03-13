package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// tokenListStep implements step.launchdarkly_token_list
type tokenListStep struct {
	name       string
	moduleName string
}

func newTokenListStep(name string, config map[string]any) (*tokenListStep, error) {
	return &tokenListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/tokens", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// tokenGetStep implements step.launchdarkly_token_get
type tokenGetStep struct {
	name       string
	moduleName string
}

func newTokenGetStep(name string, config map[string]any) (*tokenGetStep, error) {
	return &tokenGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	tokenID := resolveValue("tokenId", current, config)
	if tokenID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "tokenId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/tokens/%s", tokenID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// tokenCreateStep implements step.launchdarkly_token_create
type tokenCreateStep struct {
	name       string
	moduleName string
}

func newTokenCreateStep(name string, config map[string]any) (*tokenCreateStep, error) {
	return &tokenCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	body := map[string]any{}
	if name := resolveValue("name", current, config); name != "" {
		body["name"] = name
	}
	if role := resolveValue("role", current, config); role != "" {
		body["role"] = role
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/tokens", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// tokenUpdateStep implements step.launchdarkly_token_update
type tokenUpdateStep struct {
	name       string
	moduleName string
}

func newTokenUpdateStep(name string, config map[string]any) (*tokenUpdateStep, error) {
	return &tokenUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	tokenID := resolveValue("tokenId", current, config)
	if tokenID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "tokenId is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/tokens/%s", tokenID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// tokenDeleteStep implements step.launchdarkly_token_delete
type tokenDeleteStep struct {
	name       string
	moduleName string
}

func newTokenDeleteStep(name string, config map[string]any) (*tokenDeleteStep, error) {
	return &tokenDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	tokenID := resolveValue("tokenId", current, config)
	if tokenID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "tokenId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/tokens/%s", tokenID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// tokenResetStep implements step.launchdarkly_token_reset
type tokenResetStep struct {
	name       string
	moduleName string
}

func newTokenResetStep(name string, config map[string]any) (*tokenResetStep, error) {
	return &tokenResetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tokenResetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	tokenID := resolveValue("tokenId", current, config)
	if tokenID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "tokenId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/tokens/%s/reset", tokenID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
