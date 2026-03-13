package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// triggerListStep implements step.launchdarkly_trigger_list
type triggerListStep struct {
	name       string
	moduleName string
}

func newTriggerListStep(name string, config map[string]any) (*triggerListStep, error) {
	return &triggerListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and flagKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/flags/%s/triggers", projectKey, envKey, flagKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// triggerCreateStep implements step.launchdarkly_trigger_create
type triggerCreateStep struct {
	name       string
	moduleName string
}

func newTriggerCreateStep(name string, config map[string]any) (*triggerCreateStep, error) {
	return &triggerCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	integrationKey := resolveValue("integrationKey", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || integrationKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and integrationKey are required"}}, nil
	}
	body := map[string]any{"integrationKey": integrationKey}
	if instructions := resolveMap("instructions", current, config); instructions != nil {
		body["instructions"] = instructions
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments/%s/flags/%s/triggers", projectKey, envKey, flagKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// triggerGetStep implements step.launchdarkly_trigger_get
type triggerGetStep struct {
	name       string
	moduleName string
}

func newTriggerGetStep(name string, config map[string]any) (*triggerGetStep, error) {
	return &triggerGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	triggerID := resolveValue("triggerId", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || triggerID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and triggerId are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/flags/%s/triggers/%s", projectKey, envKey, flagKey, triggerID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// triggerUpdateStep implements step.launchdarkly_trigger_update
type triggerUpdateStep struct {
	name       string
	moduleName string
}

func newTriggerUpdateStep(name string, config map[string]any) (*triggerUpdateStep, error) {
	return &triggerUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	triggerID := resolveValue("triggerId", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || triggerID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and triggerId are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/environments/%s/flags/%s/triggers/%s", projectKey, envKey, flagKey, triggerID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// triggerDeleteStep implements step.launchdarkly_trigger_delete
type triggerDeleteStep struct {
	name       string
	moduleName string
}

func newTriggerDeleteStep(name string, config map[string]any) (*triggerDeleteStep, error) {
	return &triggerDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	triggerID := resolveValue("triggerId", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || triggerID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and triggerId are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/environments/%s/flags/%s/triggers/%s", projectKey, envKey, flagKey, triggerID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
