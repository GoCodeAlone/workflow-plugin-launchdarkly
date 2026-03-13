package internal

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// codeRefRepoListStep implements step.launchdarkly_code_ref_repo_list
type codeRefRepoListStep struct {
	name       string
	moduleName string
}

func newCodeRefRepoListStep(name string, config map[string]any) (*codeRefRepoListStep, error) {
	return &codeRefRepoListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *codeRefRepoListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, "/code-refs/repositories", nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// codeRefRepoCreateStep implements step.launchdarkly_code_ref_repo_create
type codeRefRepoCreateStep struct {
	name       string
	moduleName string
}

func newCodeRefRepoCreateStep(name string, config map[string]any) (*codeRefRepoCreateStep, error) {
	return &codeRefRepoCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *codeRefRepoCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	repoName := resolveValue("name", current, config)
	if repoName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	body := map[string]any{"name": repoName}
	if sourceLink := resolveValue("sourceLink", current, config); sourceLink != "" {
		body["sourceLink"] = sourceLink
	}
	if commitURLTemplate := resolveValue("commitUrlTemplate", current, config); commitURLTemplate != "" {
		body["commitUrlTemplate"] = commitURLTemplate
	}
	if hunkURLTemplate := resolveValue("hunkUrlTemplate", current, config); hunkURLTemplate != "" {
		body["hunkUrlTemplate"] = hunkURLTemplate
	}
	result, err := client.doRequest(ctx, http.MethodPost, "/code-refs/repositories", body)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// codeRefRepoDeleteStep implements step.launchdarkly_code_ref_repo_delete
type codeRefRepoDeleteStep struct {
	name       string
	moduleName string
}

func newCodeRefRepoDeleteStep(name string, config map[string]any) (*codeRefRepoDeleteStep, error) {
	return &codeRefRepoDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *codeRefRepoDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	repoName := resolveValue("name", current, config)
	if repoName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/code-refs/repositories/%s", repoName), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// codeRefExtinctionListStep implements step.launchdarkly_code_ref_extinction_list
type codeRefExtinctionListStep struct {
	name       string
	moduleName string
}

func newCodeRefExtinctionListStep(name string, config map[string]any) (*codeRefExtinctionListStep, error) {
	return &codeRefExtinctionListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *codeRefExtinctionListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "launchdarkly client not found: " + s.moduleName}}, nil
	}
	repoName := resolveValue("repoName", current, config)
	if repoName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "repoName is required"}}, nil
	}
	result, err := client.doRequest(ctx, http.MethodGet, fmt.Sprintf("/code-refs/repositories/%s/extinctions", repoName), nil)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
