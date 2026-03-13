package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// releasePipelineListStep implements step.launchdarkly_release_pipeline_list
type releasePipelineListStep struct {
	name       string
	moduleName string
}

func newReleasePipelineListStep(name string, config map[string]any) (*releasePipelineListStep, error) {
	return &releasePipelineListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releasePipelineListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/release-pipelines", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// releasePipelineGetStep implements step.launchdarkly_release_pipeline_get
type releasePipelineGetStep struct {
	name       string
	moduleName string
}

func newReleasePipelineGetStep(name string, config map[string]any) (*releasePipelineGetStep, error) {
	return &releasePipelineGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releasePipelineGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	pipelineKey := resolveValue("pipelineKey", current, config)
	if projectKey == "" || pipelineKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and pipelineKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/release-pipelines/%s", projectKey, pipelineKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// releasePipelineCreateStep implements step.launchdarkly_release_pipeline_create
type releasePipelineCreateStep struct {
	name       string
	moduleName string
}

func newReleasePipelineCreateStep(name string, config map[string]any) (*releasePipelineCreateStep, error) {
	return &releasePipelineCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releasePipelineCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	pipelineKey := resolveValue("key", current, config)
	pipelineName := resolveValue("name", current, config)
	if projectKey == "" || pipelineKey == "" || pipelineName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, key, and name are required"}}, nil
	}
	body := map[string]any{
		"key":  pipelineKey,
		"name": pipelineName,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/release-pipelines", projectKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// releasePipelineUpdateStep implements step.launchdarkly_release_pipeline_update
type releasePipelineUpdateStep struct {
	name       string
	moduleName string
}

func newReleasePipelineUpdateStep(name string, config map[string]any) (*releasePipelineUpdateStep, error) {
	return &releasePipelineUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releasePipelineUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	pipelineKey := resolveValue("pipelineKey", current, config)
	if projectKey == "" || pipelineKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and pipelineKey are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/release-pipelines/%s", projectKey, pipelineKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// releasePipelineDeleteStep implements step.launchdarkly_release_pipeline_delete
type releasePipelineDeleteStep struct {
	name       string
	moduleName string
}

func newReleasePipelineDeleteStep(name string, config map[string]any) (*releasePipelineDeleteStep, error) {
	return &releasePipelineDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releasePipelineDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	pipelineKey := resolveValue("pipelineKey", current, config)
	if projectKey == "" || pipelineKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and pipelineKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/release-pipelines/%s", projectKey, pipelineKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
