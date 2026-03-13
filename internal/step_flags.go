package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// flagListStep implements step.launchdarkly_flag_list
type flagListStep struct {
	name       string
	moduleName string
}

func newFlagListStep(name string, config map[string]any) (*flagListStep, error) {
	return &flagListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/flags/%s", projectKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagGetStep implements step.launchdarkly_flag_get
type flagGetStep struct {
	name       string
	moduleName string
}

func newFlagGetStep(name string, config map[string]any) (*flagGetStep, error) {
	return &flagGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	if flagKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flagKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/flags/%s/%s", projectKey, flagKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagCreateStep implements step.launchdarkly_flag_create
type flagCreateStep struct {
	name       string
	moduleName string
}

func newFlagCreateStep(name string, config map[string]any) (*flagCreateStep, error) {
	return &flagCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	flagName := resolveValue("name", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	if flagKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flagKey is required"}}, nil
	}
	if flagName == "" {
		flagName = flagKey
	}
	body := map[string]any{
		"key":  flagKey,
		"name": flagName,
	}
	if kind := resolveValue("kind", current, config); kind != "" {
		body["kind"] = kind
	} else {
		body["kind"] = "boolean"
	}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	if tags := resolveStringSlice("tags", current, config); len(tags) > 0 {
		tagsAny := make([]any, len(tags))
		for i, t := range tags {
			tagsAny[i] = t
		}
		body["tags"] = tagsAny
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/flags/%s", projectKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagUpdateStep implements step.launchdarkly_flag_update
type flagUpdateStep struct {
	name       string
	moduleName string
}

func newFlagUpdateStep(name string, config map[string]any) (*flagUpdateStep, error) {
	return &flagUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	if flagKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flagKey is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/flags/%s/%s", projectKey, flagKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagDeleteStep implements step.launchdarkly_flag_delete
type flagDeleteStep struct {
	name       string
	moduleName string
}

func newFlagDeleteStep(name string, config map[string]any) (*flagDeleteStep, error) {
	return &flagDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	if projectKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey is required"}}, nil
	}
	if flagKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flagKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/flags/%s/%s", projectKey, flagKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagCopyStep implements step.launchdarkly_flag_copy
type flagCopyStep struct {
	name       string
	moduleName string
}

func newFlagCopyStep(name string, config map[string]any) (*flagCopyStep, error) {
	return &flagCopyStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagCopyStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	flagKey := resolveValue("flagKey", current, config)
	sourceEnv := resolveValue("sourceEnvironment", current, config)
	targetEnv := resolveValue("targetEnvironment", current, config)
	if projectKey == "" || flagKey == "" || sourceEnv == "" || targetEnv == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey, flagKey, sourceEnvironment, and targetEnvironment are required"}}, nil
	}
	body := map[string]any{
		"source": map[string]any{"key": sourceEnv},
		"target": map[string]any{"key": targetEnv},
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/flags/%s/%s/copy", projectKey, flagKey), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagStatusGetStep implements step.launchdarkly_flag_status_get
type flagStatusGetStep struct {
	name       string
	moduleName string
}

func newFlagStatusGetStep(name string, config map[string]any) (*flagStatusGetStep, error) {
	return &flagStatusGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagStatusGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
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
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/flag-statuses/%s/%s/%s", projectKey, envKey, flagKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// flagStatusListStep implements step.launchdarkly_flag_status_list
type flagStatusListStep struct {
	name       string
	moduleName string
}

func newFlagStatusListStep(name string, config map[string]any) (*flagStatusListStep, error) {
	return &flagStatusListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *flagStatusListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	projectKey := resolveValue("projectKey", current, config)
	envKey := resolveValue("environmentKey", current, config)
	if projectKey == "" || envKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "projectKey and environmentKey are required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/flag-statuses/%s/%s", projectKey, envKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
