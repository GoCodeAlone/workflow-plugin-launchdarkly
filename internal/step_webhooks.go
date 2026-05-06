package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// webhookListStep implements step.launchdarkly_webhook_list
type webhookListStep struct {
	name       string
	moduleName string
}

func newWebhookListStep(name string, config map[string]any) (*webhookListStep, error) {
	return &webhookListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webhookListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/webhooks", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// webhookGetStep implements step.launchdarkly_webhook_get
type webhookGetStep struct {
	name       string
	moduleName string
}

func newWebhookGetStep(name string, config map[string]any) (*webhookGetStep, error) {
	return &webhookGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webhookGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	webhookID := resolveValue("webhookId", current, config)
	if webhookID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "webhookId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/webhooks/%s", webhookID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// webhookCreateStep implements step.launchdarkly_webhook_create
type webhookCreateStep struct {
	name       string
	moduleName string
}

func newWebhookCreateStep(name string, config map[string]any) (*webhookCreateStep, error) {
	return &webhookCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webhookCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	url := resolveValue("url", current, config)
	if url == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "url is required"}}, nil
	}
	body := map[string]any{"url": url}
	if v, ok := current["sign"]; ok {
		body["sign"] = v
	} else if v, ok := config["sign"]; ok {
		body["sign"] = v
	}
	if v, ok := current["on"]; ok {
		body["on"] = v
	} else if v, ok := config["on"]; ok {
		body["on"] = v
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/webhooks", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// webhookUpdateStep implements step.launchdarkly_webhook_update
type webhookUpdateStep struct {
	name       string
	moduleName string
}

func newWebhookUpdateStep(name string, config map[string]any) (*webhookUpdateStep, error) {
	return &webhookUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webhookUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	webhookID := resolveValue("webhookId", current, config)
	if webhookID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "webhookId is required"}}, nil
	}
	patch := resolveMap("patch", current, config)
	if patch == nil {
		patch = map[string]any{}
	}
	result, err := client.doRequest(ctx, http.MethodPatch, fmt.Sprintf("/webhooks/%s", webhookID), patch)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// webhookDeleteStep implements step.launchdarkly_webhook_delete
type webhookDeleteStep struct {
	name       string
	moduleName string
}

func newWebhookDeleteStep(name string, config map[string]any) (*webhookDeleteStep, error) {
	return &webhookDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webhookDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	webhookID := resolveValue("webhookId", current, config)
	if webhookID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "webhookId is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/webhooks/%s", webhookID), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
