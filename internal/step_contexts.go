package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// contextListStep implements step.launchdarkly_context_list
type contextListStep struct {
	name       string
	moduleName string
}

func newContextListStep(name string, config map[string]any) (*contextListStep, error) {
	return &contextListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/contexts", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// contextGetStep implements step.launchdarkly_context_get
type contextGetStep struct {
	name       string
	moduleName string
}

func newContextGetStep(name string, config map[string]any) (*contextGetStep, error) {
	return &contextGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	kind := resolveValue("kind", current, config)
	contextKey := resolveValue("contextKey", current, config)
	if projectKey == "" || envKey == "" || kind == "" || contextKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, kind, and contextKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/contexts/%s/%s", projectKey, envKey, kind, contextKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// contextSearchStep implements step.launchdarkly_context_search
type contextSearchStep struct {
	name       string
	moduleName string
}

func newContextSearchStep(name string, config map[string]any) (*contextSearchStep, error) {
	return &contextSearchStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextSearchStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	body := map[string]any{}
	if filter := resolveValue("filter", current, config); filter != "" {
		body["filter"] = filter
	}
	if sort := resolveValue("sort", current, config); sort != "" {
		body["sort"] = sort
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments/%s/contexts/search", projectKey, envKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// contextKindListStep implements step.launchdarkly_context_kind_list
type contextKindListStep struct {
	name       string
	moduleName string
}

func newContextKindListStep(name string, config map[string]any) (*contextKindListStep, error) {
	return &contextKindListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextKindListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/context-kinds", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// contextKindUpsertStep implements step.launchdarkly_context_kind_upsert
type contextKindUpsertStep struct {
	name       string
	moduleName string
}

func newContextKindUpsertStep(name string, config map[string]any) (*contextKindUpsertStep, error) {
	return &contextKindUpsertStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextKindUpsertStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	kind := resolveValue("key", current, config)
	kindName := resolveValue("name", current, config)
	if projectKey == "" || kind == "" || kindName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, key, and name are required"}}, nil
	}
	body := map[string]any{"name": kindName}
	result, err := client.doRequest(ctx, http.MethodPut, fmt.Sprintf("/projects/%s/context-kinds/%s", projectKey, kind), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// contextEvaluateStep implements step.launchdarkly_context_evaluate
type contextEvaluateStep struct {
	name       string
	moduleName string
}

func newContextEvaluateStep(name string, config map[string]any) (*contextEvaluateStep, error) {
	return &contextEvaluateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *contextEvaluateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	kind := resolveValue("kind", current, config)
	contextKey := resolveValue("contextKey", current, config)
	if projectKey == "" || envKey == "" || kind == "" || contextKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, kind, and contextKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/flags/evaluate/%s/%s", projectKey, envKey, kind, contextKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
