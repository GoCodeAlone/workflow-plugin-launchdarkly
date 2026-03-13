package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// experimentListStep implements step.launchdarkly_experiment_list
type experimentListStep struct {
	name       string
	moduleName string
}

func newExperimentListStep(name string, config map[string]any) (*experimentListStep, error) {
	return &experimentListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *experimentListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/experiments", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// experimentGetStep implements step.launchdarkly_experiment_get
type experimentGetStep struct {
	name       string
	moduleName string
}

func newExperimentGetStep(name string, config map[string]any) (*experimentGetStep, error) {
	return &experimentGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *experimentGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	experimentKey := resolveValue("experimentKey", current, config)
	if projectKey == "" || envKey == "" || experimentKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and experimentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/experiments/%s", projectKey, envKey, experimentKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// experimentCreateStep implements step.launchdarkly_experiment_create
type experimentCreateStep struct {
	name       string
	moduleName string
}

func newExperimentCreateStep(name string, config map[string]any) (*experimentCreateStep, error) {
	return &experimentCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *experimentCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	experimentName := resolveValue("name", current, config)
	if projectKey == "" || envKey == "" || experimentName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and name are required"}}, nil
	}
	body := map[string]any{"name": experimentName}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	if iteration := resolveMap("iteration", current, config); iteration != nil {
		body["iteration"] = iteration
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/environments/%s/experiments", projectKey, envKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// experimentUpdateStep implements step.launchdarkly_experiment_update
type experimentUpdateStep struct {
	name       string
	moduleName string
}

func newExperimentUpdateStep(name string, config map[string]any) (*experimentUpdateStep, error) {
	return &experimentUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *experimentUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	experimentKey := resolveValue("experimentKey", current, config)
	if projectKey == "" || envKey == "" || experimentKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, and experimentKey are required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/environments/%s/experiments/%s", projectKey, envKey, experimentKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// experimentResultsGetStep implements step.launchdarkly_experiment_results_get
type experimentResultsGetStep struct {
	name       string
	moduleName string
}

func newExperimentResultsGetStep(name string, config map[string]any) (*experimentResultsGetStep, error) {
	return &experimentResultsGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *experimentResultsGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	experimentKey := resolveValue("experimentKey", current, config)
	metricKey := resolveValue("metricKey", current, config)
	if projectKey == "" || envKey == "" || experimentKey == "" || metricKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, environmentKey, experimentKey, and metricKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/environments/%s/experiments/%s/metrics/%s/results", projectKey, envKey, experimentKey, metricKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
