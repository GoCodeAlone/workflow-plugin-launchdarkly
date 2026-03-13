package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// auditLogListStep implements step.launchdarkly_audit_log_list
type auditLogListStep struct {
	name       string
	moduleName string
}

func newAuditLogListStep(name string, config map[string]any) (*auditLogListStep, error) {
	return &auditLogListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *auditLogListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/auditlog", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// auditLogGetStep implements step.launchdarkly_audit_log_get
type auditLogGetStep struct {
	name       string
	moduleName string
}

func newAuditLogGetStep(name string, config map[string]any) (*auditLogGetStep, error) {
	return &auditLogGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *auditLogGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	specID := resolveValue("specId", current, config)
	if specID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "specId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/auditlog/%s", specID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
