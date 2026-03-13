package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// workflowListStep implements step.launchdarkly_workflow_list
type workflowListStep struct {
	name       string
	moduleName string
}

func newWorkflowListStep(name string, config map[string]any) (*workflowListStep, error) {
	return &workflowListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *workflowListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
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
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/flags/%s/environments/%s/workflows", projectKey, flagKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// workflowGetStep implements step.launchdarkly_workflow_get
type workflowGetStep struct {
	name       string
	moduleName string
}

func newWorkflowGetStep(name string, config map[string]any) (*workflowGetStep, error) {
	return &workflowGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *workflowGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	workflowID := resolveValue("workflowId", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || workflowID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and workflowId are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/flags/%s/environments/%s/workflows/%s", projectKey, flagKey, envKey, workflowID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// workflowCreateStep implements step.launchdarkly_workflow_create
type workflowCreateStep struct {
	name       string
	moduleName string
}

func newWorkflowCreateStep(name string, config map[string]any) (*workflowCreateStep, error) {
	return &workflowCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *workflowCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	workflowName := resolveValue("name", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || workflowName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and name are required"}}, nil
	}
	body := map[string]any{"name": workflowName}
	if stages := resolveMap("stages", current, config); stages != nil {
		body["stages"] = stages
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/flags/%s/environments/%s/workflows", projectKey, flagKey, envKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// workflowDeleteStep implements step.launchdarkly_workflow_delete
type workflowDeleteStep struct {
	name       string
	moduleName string
}

func newWorkflowDeleteStep(name string, config map[string]any) (*workflowDeleteStep, error) {
	return &workflowDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *workflowDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	workflowID := resolveValue("workflowId", current, config)
	if projectKey == "" || envKey == "" || flagKey == "" || workflowID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, flagKey, and workflowId are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/flags/%s/environments/%s/workflows/%s", projectKey, flagKey, envKey, workflowID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
