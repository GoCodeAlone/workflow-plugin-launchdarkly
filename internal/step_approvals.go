package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// approvalListStep implements step.launchdarkly_approval_list
type approvalListStep struct {
	name       string
	moduleName string
}

func newApprovalListStep(name string, config map[string]any) (*approvalListStep, error) {
	return &approvalListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/approval-requests", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// approvalGetStep implements step.launchdarkly_approval_get
type approvalGetStep struct {
	name       string
	moduleName string
}

func newApprovalGetStep(name string, config map[string]any) (*approvalGetStep, error) {
	return &approvalGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	approvalID := resolveValue("approvalId", current, config)
	if approvalID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "approvalId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/approval-requests/%s", approvalID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// approvalCreateStep implements step.launchdarkly_approval_create
type approvalCreateStep struct {
	name       string
	moduleName string
}

func newApprovalCreateStep(name string, config map[string]any) (*approvalCreateStep, error) {
	return &approvalCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	description := resolveValue("description", current, config)
	if description == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "description is required"}}, nil
	}
	body := map[string]any{"description": description}
	if instructions := resolveMap("instructions", current, config); instructions != nil {
		body["instructions"] = instructions
	}
	if notifyMemberIDs := resolveStringSlice("notifyMemberIds", current, config); len(notifyMemberIDs) > 0 {
		ids := make([]any, len(notifyMemberIDs))
		for i, id := range notifyMemberIDs {
			ids[i] = id
		}
		body["notifyMemberIds"] = ids
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/approval-requests", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// approvalDeleteStep implements step.launchdarkly_approval_delete
type approvalDeleteStep struct {
	name       string
	moduleName string
}

func newApprovalDeleteStep(name string, config map[string]any) (*approvalDeleteStep, error) {
	return &approvalDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	approvalID := resolveValue("approvalId", current, config)
	if approvalID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "approvalId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/approval-requests/%s", approvalID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// approvalApplyStep implements step.launchdarkly_approval_apply
type approvalApplyStep struct {
	name       string
	moduleName string
}

func newApprovalApplyStep(name string, config map[string]any) (*approvalApplyStep, error) {
	return &approvalApplyStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalApplyStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	approvalID := resolveValue("approvalId", current, config)
	if approvalID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "approvalId is required"}}, nil
	}
	body := map[string]any{}
	if comment := resolveValue("comment", current, config); comment != "" {
		body["comment"] = comment
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/approval-requests/%s/apply", approvalID), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// approvalReviewStep implements step.launchdarkly_approval_review
type approvalReviewStep struct {
	name       string
	moduleName string
}

func newApprovalReviewStep(name string, config map[string]any) (*approvalReviewStep, error) {
	return &approvalReviewStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *approvalReviewStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	approvalID := resolveValue("approvalId", current, config)
	kind := resolveValue("kind", current, config)
	if approvalID == "" || kind == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "approvalId and kind are required"}}, nil
	}
	body := map[string]any{"kind": kind}
	if comment := resolveValue("comment", current, config); comment != "" {
		body["comment"] = comment
	}
	result, err := client.doRequest(ctx, http.MethodPost, fmt.Sprintf("/approval-requests/%s/reviews", approvalID), body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
