package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// projectListStep implements step.launchdarkly_project_list
type projectListStep struct {
	name       string
	moduleName string
}

func newProjectListStep(name string, config map[string]any) (*projectListStep, error) {
	return &projectListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *projectListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/projects", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// projectGetStep implements step.launchdarkly_project_get
type projectGetStep struct {
	name       string
	moduleName string
}

func newProjectGetStep(name string, config map[string]any) (*projectGetStep, error) {
	return &projectGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *projectGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// projectCreateStep implements step.launchdarkly_project_create
type projectCreateStep struct {
	name       string
	moduleName string
}

func newProjectCreateStep(name string, config map[string]any) (*projectCreateStep, error) {
	return &projectCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *projectCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	projectName := resolveValue("name", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	if projectName == "" {
		projectName = projectKey
	}
	body := map[string]any{
		"key":  projectKey,
		"name": projectName,
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/projects", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// projectUpdateStep implements step.launchdarkly_project_update
type projectUpdateStep struct {
	name       string
	moduleName string
}

func newProjectUpdateStep(name string, config map[string]any) (*projectUpdateStep, error) {
	return &projectUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *projectUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s", projectKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// projectDeleteStep implements step.launchdarkly_project_delete
type projectDeleteStep struct {
	name       string
	moduleName string
}

func newProjectDeleteStep(name string, config map[string]any) (*projectDeleteStep, error) {
	return &projectDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *projectDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
