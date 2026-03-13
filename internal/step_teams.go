package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// teamListStep implements step.launchdarkly_team_list
type teamListStep struct {
	name       string
	moduleName string
}

func newTeamListStep(name string, config map[string]any) (*teamListStep, error) {
	return &teamListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *teamListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/teams", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// teamGetStep implements step.launchdarkly_team_get
type teamGetStep struct {
	name       string
	moduleName string
}

func newTeamGetStep(name string, config map[string]any) (*teamGetStep, error) {
	return &teamGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *teamGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	teamKey := resolveValue("teamKey", current, config)
	if teamKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "teamKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/teams/%s", teamKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// teamCreateStep implements step.launchdarkly_team_create
type teamCreateStep struct {
	name       string
	moduleName string
}

func newTeamCreateStep(name string, config map[string]any) (*teamCreateStep, error) {
	return &teamCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *teamCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	teamKey := resolveValue("teamKey", current, config)
	teamName := resolveValue("name", current, config)
	if teamKey == "" || teamName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "teamKey and name are required"}}, nil
	}
	body := map[string]any{
		"key":  teamKey,
		"name": teamName,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		body["description"] = desc
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/teams", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// teamUpdateStep implements step.launchdarkly_team_update
type teamUpdateStep struct {
	name       string
	moduleName string
}

func newTeamUpdateStep(name string, config map[string]any) (*teamUpdateStep, error) {
	return &teamUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *teamUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	teamKey := resolveValue("teamKey", current, config)
	if teamKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "teamKey is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/teams/%s", teamKey), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// teamDeleteStep implements step.launchdarkly_team_delete
type teamDeleteStep struct {
	name       string
	moduleName string
}

func newTeamDeleteStep(name string, config map[string]any) (*teamDeleteStep, error) {
	return &teamDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *teamDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	teamKey := resolveValue("teamKey", current, config)
	if teamKey == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "teamKey is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/teams/%s", teamKey), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
