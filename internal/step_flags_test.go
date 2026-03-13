package internal

import (
	"context"
	"testing"
)

func TestFlagListStep_MissingClient(t *testing.T) {
	step, err := newFlagListStep("test", map[string]any{"module": "nonexistent-flags"})
	if err != nil {
		t.Fatal(err)
	}
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{"projectKey": "my-project"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFlagListStep_MissingProjectKey(t *testing.T) {
	step, _ := newFlagListStep("test", map[string]any{"module": "nonexistent-flags2"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing projectKey")
	}
}

func TestFlagGetStep_MissingFlagKey(t *testing.T) {
	step, _ := newFlagGetStep("test", map[string]any{"module": "nonexistent-flag-get"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{"projectKey": "my-project"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing flagKey")
	}
}

func TestFlagCreateStep_MissingProjectKey(t *testing.T) {
	step, _ := newFlagCreateStep("test", map[string]any{"module": "nonexistent-flag-create"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing projectKey")
	}
}

func TestFlagDeleteStep_MissingClient(t *testing.T) {
	step, _ := newFlagDeleteStep("test", map[string]any{"module": "nonexistent-flag-del"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{"projectKey": "proj", "flagKey": "flag"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFlagCopyStep_MissingParams(t *testing.T) {
	step, _ := newFlagCopyStep("test", map[string]any{"module": "nonexistent-flag-copy"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing params")
	}
}

func TestFlagStatusGetStep_MissingParams(t *testing.T) {
	step, _ := newFlagStatusGetStep("test", map[string]any{"module": "nonexistent-flag-status"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing params")
	}
}

func TestStepRegistry_AllTypesInstantiable(t *testing.T) {
	for typeName, constructor := range stepRegistry {
		step, err := constructor("test-"+typeName, map[string]any{})
		if err != nil {
			t.Errorf("step type %q constructor failed: %v", typeName, err)
			continue
		}
		if step == nil {
			t.Errorf("step type %q returned nil", typeName)
		}
	}
}

func TestAllStepTypes_NonEmpty(t *testing.T) {
	types := allStepTypes()
	if len(types) == 0 {
		t.Error("expected non-empty step types")
	}
	// We expect ~100 step types
	if len(types) < 90 {
		t.Errorf("expected at least 90 step types, got %d", len(types))
	}
}

func TestCreateStep_UnknownType(t *testing.T) {
	_, err := createStep("step.nonexistent_type", "test", map[string]any{})
	if err == nil {
		t.Error("expected error for unknown step type")
	}
}
