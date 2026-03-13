package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// memberListStep implements step.launchdarkly_member_list
type memberListStep struct {
	name       string
	moduleName string
}

func newMemberListStep(name string, config map[string]any) (*memberListStep, error) {
	return &memberListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *memberListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/members", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// memberGetStep implements step.launchdarkly_member_get
type memberGetStep struct {
	name       string
	moduleName string
}

func newMemberGetStep(name string, config map[string]any) (*memberGetStep, error) {
	return &memberGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *memberGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	memberID := resolveValue("memberId", current, config)
	if memberID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "memberId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/members/%s", memberID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// memberCreateStep implements step.launchdarkly_member_create
type memberCreateStep struct {
	name       string
	moduleName string
}

func newMemberCreateStep(name string, config map[string]any) (*memberCreateStep, error) {
	return &memberCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *memberCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	email := resolveValue("email", current, config)
	if email == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "email is required"}}, nil
	}
	body := []any{map[string]any{"email": email}}
	if role := resolveValue("role", current, config); role != "" {
		body[0].(map[string]any)["role"] = role
	}
	if firstName := resolveValue("firstName", current, config); firstName != "" {
		body[0].(map[string]any)["firstName"] = firstName
	}
	if lastName := resolveValue("lastName", current, config); lastName != "" {
		body[0].(map[string]any)["lastName"] = lastName
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/members", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// memberUpdateStep implements step.launchdarkly_member_update
type memberUpdateStep struct {
	name       string
	moduleName string
}

func newMemberUpdateStep(name string, config map[string]any) (*memberUpdateStep, error) {
	return &memberUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *memberUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	memberID := resolveValue("memberId", current, config)
	if memberID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "memberId is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/members/%s", memberID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// memberDeleteStep implements step.launchdarkly_member_delete
type memberDeleteStep struct {
	name       string
	moduleName string
}

func newMemberDeleteStep(name string, config map[string]any) (*memberDeleteStep, error) {
	return &memberDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *memberDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	memberID := resolveValue("memberId", current, config)
	if memberID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "memberId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/members/%s", memberID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
