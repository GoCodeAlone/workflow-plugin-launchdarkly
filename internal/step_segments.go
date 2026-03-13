package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// segmentListStep implements step.launchdarkly_segment_list
type segmentListStep struct {
	name       string
	moduleName string
}

func newSegmentListStep(name string, config map[string]any) (*segmentListStep, error) {
	return &segmentListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *segmentListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/segments/%s/%s", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// segmentGetStep implements step.launchdarkly_segment_get
type segmentGetStep struct {
	name       string
	moduleName string
}

func newSegmentGetStep(name string, config map[string]any) (*segmentGetStep, error) {
	return &segmentGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *segmentGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	segmentKey := resolveValue("segmentKey", current, config)
	if projectKey == "" || envKey == "" || segmentKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and segmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/segments/%s/%s/%s", projectKey, envKey, segmentKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// segmentCreateStep implements step.launchdarkly_segment_create
type segmentCreateStep struct {
	name       string
	moduleName string
}

func newSegmentCreateStep(name string, config map[string]any) (*segmentCreateStep, error) {
	return &segmentCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *segmentCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	segmentKey := resolveValue("segmentKey", current, config)
	segmentName := resolveValue("name", current, config)
	if projectKey == "" || envKey == "" || segmentKey == "" || segmentName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, segmentKey, and name are required"}}, nil
	}
	body := map[string]any{
		"key":  segmentKey,
		"name": segmentName,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/segments/%s/%s", projectKey, envKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// segmentUpdateStep implements step.launchdarkly_segment_update
type segmentUpdateStep struct {
	name       string
	moduleName string
}

func newSegmentUpdateStep(name string, config map[string]any) (*segmentUpdateStep, error) {
	return &segmentUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *segmentUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	segmentKey := resolveValue("segmentKey", current, config)
	if projectKey == "" || envKey == "" || segmentKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and segmentKey are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/segments/%s/%s/%s", projectKey, envKey, segmentKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// segmentDeleteStep implements step.launchdarkly_segment_delete
type segmentDeleteStep struct {
	name       string
	moduleName string
}

func newSegmentDeleteStep(name string, config map[string]any) (*segmentDeleteStep, error) {
	return &segmentDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *segmentDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	segmentKey := resolveValue("segmentKey", current, config)
	if projectKey == "" || envKey == "" || segmentKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and segmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/segments/%s/%s/%s", projectKey, envKey, segmentKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
