package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// roleListStep implements step.launchdarkly_role_list
type roleListStep struct {
	name       string
	moduleName string
}

func newRoleListStep(name string, config map[string]any) (*roleListStep, error) {
	return &roleListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/roles", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// roleGetStep implements step.launchdarkly_role_get
type roleGetStep struct {
	name       string
	moduleName string
}

func newRoleGetStep(name string, config map[string]any) (*roleGetStep, error) {
	return &roleGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	customRoleKey := resolveValue("customRoleKey", current, config)
	if customRoleKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "customRoleKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/roles/%s", customRoleKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// roleCreateStep implements step.launchdarkly_role_create
type roleCreateStep struct {
	name       string
	moduleName string
}

func newRoleCreateStep(name string, config map[string]any) (*roleCreateStep, error) {
	return &roleCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	roleKey := resolveValue("key", current, config)
	roleName := resolveValue("name", current, config)
	if roleKey == "" || roleName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "key and name are required"}}, nil
	}
	body := map[string]any{
		"key":  roleKey,
		"name": roleName,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/roles", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// roleUpdateStep implements step.launchdarkly_role_update
type roleUpdateStep struct {
	name       string
	moduleName string
}

func newRoleUpdateStep(name string, config map[string]any) (*roleUpdateStep, error) {
	return &roleUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	customRoleKey := resolveValue("customRoleKey", current, config)
	if customRoleKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "customRoleKey is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/roles/%s", customRoleKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// roleDeleteStep implements step.launchdarkly_role_delete
type roleDeleteStep struct {
	name       string
	moduleName string
}

func newRoleDeleteStep(name string, config map[string]any) (*roleDeleteStep, error) {
	return &roleDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	customRoleKey := resolveValue("customRoleKey", current, config)
	if customRoleKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "customRoleKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/roles/%s", customRoleKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
