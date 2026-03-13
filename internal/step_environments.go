package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// environmentListStep implements step.launchdarkly_environment_list
type environmentListStep struct {
	name       string
	moduleName string
}

func newEnvironmentListStep(name string, config map[string]any) (*environmentListStep, error) {
	return &environmentListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentGetStep implements step.launchdarkly_environment_get
type environmentGetStep struct {
	name       string
	moduleName string
}

func newEnvironmentGetStep(name string, config map[string]any) (*environmentGetStep, error) {
	return &environmentGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentCreateStep implements step.launchdarkly_environment_create
type environmentCreateStep struct {
	name       string
	moduleName string
}

func newEnvironmentCreateStep(name string, config map[string]any) (*environmentCreateStep, error) {
	return &environmentCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	envName := resolveValue("name", current, config)
	color := resolveValue("color", current, config)
	if projectKey == "" || envKey == "" || envName == "" || color == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, name, and color are required"}}, nil
	}
	body := map[string]any{
		"key":   envKey,
		"name":  envName,
		"color": color,
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments", projectKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentUpdateStep implements step.launchdarkly_environment_update
type environmentUpdateStep struct {
	name       string
	moduleName string
}

func newEnvironmentUpdateStep(name string, config map[string]any) (*environmentUpdateStep, error) {
	return &environmentUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/environments/%s", projectKey, envKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentDeleteStep implements step.launchdarkly_environment_delete
type environmentDeleteStep struct {
	name       string
	moduleName string
}

func newEnvironmentDeleteStep(name string, config map[string]any) (*environmentDeleteStep, error) {
	return &environmentDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/environments/%s", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentResetSDKKeyStep implements step.launchdarkly_environment_reset_sdk_key
type environmentResetSDKKeyStep struct {
	name       string
	moduleName string
}

func newEnvironmentResetSDKKeyStep(name string, config map[string]any) (*environmentResetSDKKeyStep, error) {
	return &environmentResetSDKKeyStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentResetSDKKeyStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments/%s/apiKey", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// environmentResetMobileKeyStep implements step.launchdarkly_environment_reset_mobile_key
type environmentResetMobileKeyStep struct {
	name       string
	moduleName string
}

func newEnvironmentResetMobileKeyStep(name string, config map[string]any) (*environmentResetMobileKeyStep, error) {
	return &environmentResetMobileKeyStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *environmentResetMobileKeyStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments/%s/mobileKey", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
