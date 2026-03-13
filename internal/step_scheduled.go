package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// scheduledChangeListStep implements step.launchdarkly_scheduled_change_list
type scheduledChangeListStep struct {
	name       string
	moduleName string
}

func newScheduledChangeListStep(name string, config map[string]any) (*scheduledChangeListStep, error) {
	return &scheduledChangeListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scheduledChangeListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
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
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/flags/%s/scheduled-changes", projectKey, flagKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	_ = envKey
	return &sdk.StepResult{Output: result}, nil
}

// scheduledChangeCreateStep implements step.launchdarkly_scheduled_change_create
type scheduledChangeCreateStep struct {
	name       string
	moduleName string
}

func newScheduledChangeCreateStep(name string, config map[string]any) (*scheduledChangeCreateStep, error) {
	return &scheduledChangeCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scheduledChangeCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	executionDate := resolveValue("executionDate", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || executionDate == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and executionDate are required"}}, nil
	}
	body := map[string]any{
		"environmentKey": envKey,
		"executionDate":  executionDate,
	}
	if instructions := resolveMap("instructions", current, config); instructions != nil {
		body["instructions"] = instructions
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/flags/%s/scheduled-changes", projectKey, flagKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// scheduledChangeUpdateStep implements step.launchdarkly_scheduled_change_update
type scheduledChangeUpdateStep struct {
	name       string
	moduleName string
}

func newScheduledChangeUpdateStep(name string, config map[string]any) (*scheduledChangeUpdateStep, error) {
	return &scheduledChangeUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scheduledChangeUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	scheduledChangeID := resolveValue("scheduledChangeId", current, config)
	if projectKey == "" || flagKey == "" || scheduledChangeID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, flagKey, and scheduledChangeId are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/flags/%s/scheduled-changes/%s", projectKey, flagKey, scheduledChangeID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// scheduledChangeDeleteStep implements step.launchdarkly_scheduled_change_delete
type scheduledChangeDeleteStep struct {
	name       string
	moduleName string
}

func newScheduledChangeDeleteStep(name string, config map[string]any) (*scheduledChangeDeleteStep, error) {
	return &scheduledChangeDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scheduledChangeDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	scheduledChangeID := resolveValue("scheduledChangeId", current, config)
	if projectKey == "" || flagKey == "" || scheduledChangeID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, flagKey, and scheduledChangeId are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/flags/%s/scheduled-changes/%s", projectKey, flagKey, scheduledChangeID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
