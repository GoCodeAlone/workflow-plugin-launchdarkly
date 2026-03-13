package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// metricListStep implements step.launchdarkly_metric_list
type metricListStep struct {
	name       string
	moduleName string
}

func newMetricListStep(name string, config map[string]any) (*metricListStep, error) {
	return &metricListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *metricListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/metrics/%s", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// metricGetStep implements step.launchdarkly_metric_get
type metricGetStep struct {
	name       string
	moduleName string
}

func newMetricGetStep(name string, config map[string]any) (*metricGetStep, error) {
	return &metricGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *metricGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	metricKey := resolveValue("metricKey", current, config)
	if projectKey == "" || metricKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and metricKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/metrics/%s/%s", projectKey, metricKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// metricCreateStep implements step.launchdarkly_metric_create
type metricCreateStep struct {
	name       string
	moduleName string
}

func newMetricCreateStep(name string, config map[string]any) (*metricCreateStep, error) {
	return &metricCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *metricCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	metricKey := resolveValue("metricKey", current, config)
	metricName := resolveValue("name", current, config)
	kind := resolveValue("kind", current, config)
	if projectKey == "" || metricKey == "" || metricName == "" || kind == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, metricKey, name, and kind are required"}}, nil
	}
	body := map[string]any{
		"key":  metricKey,
		"name": metricName,
		"kind": kind,
	}
	if eventKey := resolveValue("eventKey", current, config); eventKey != "" {
		body["eventKey"] = eventKey
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/metrics/%s", projectKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// metricUpdateStep implements step.launchdarkly_metric_update
type metricUpdateStep struct {
	name       string
	moduleName string
}

func newMetricUpdateStep(name string, config map[string]any) (*metricUpdateStep, error) {
	return &metricUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *metricUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	metricKey := resolveValue("metricKey", current, config)
	if projectKey == "" || metricKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and metricKey are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/metrics/%s/%s", projectKey, metricKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// metricDeleteStep implements step.launchdarkly_metric_delete
type metricDeleteStep struct {
	name       string
	moduleName string
}

func newMetricDeleteStep(name string, config map[string]any) (*metricDeleteStep, error) {
	return &metricDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *metricDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	metricKey := resolveValue("metricKey", current, config)
	if projectKey == "" || metricKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and metricKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/metrics/%s/%s", projectKey, metricKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
