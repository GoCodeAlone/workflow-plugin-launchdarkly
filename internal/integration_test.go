package internal_test

import (
	"testing"

	"github.com/GoCodeAlone/workflow/wftest"
)

func TestIntegration_FlagList(t *testing.T) {
	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  list_flags:
    steps:
      - name: list
        type: step.launchdarkly_flag_list
        config:
          projectKey: "my-project"
      - name: mark_done
        type: step.set
        config:
          values:
            done: true
`),
		wftest.MockStep("step.launchdarkly_flag_list", wftest.Returns(map[string]any{
			"items": []any{
				map[string]any{"key": "flag-a", "name": "Flag A"},
				map[string]any{"key": "flag-b", "name": "Flag B"},
			},
			"totalCount": 2,
		})),
	)

	result := h.ExecutePipeline("list_flags", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["done"] != true {
		t.Error("expected done=true")
	}
	flagItems := result.StepOutput("list")
	if flagItems == nil {
		t.Fatal("expected step output from 'list' step")
	}
	if flagItems["totalCount"] != 2 {
		t.Errorf("expected totalCount=2, got %v", flagItems["totalCount"])
	}
}

func TestIntegration_FlagGet(t *testing.T) {
	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  get_flag:
    steps:
      - name: get
        type: step.launchdarkly_flag_get
        config:
          projectKey: "my-project"
          flagKey: "my-feature"
      - name: mark_done
        type: step.set
        config:
          values:
            fetched: true
`),
		wftest.MockStep("step.launchdarkly_flag_get", wftest.Returns(map[string]any{
			"key":    "my-feature",
			"name":   "My Feature",
			"kind":   "boolean",
			"on":     true,
		})),
	)

	result := h.ExecutePipeline("get_flag", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["fetched"] != true {
		t.Error("expected fetched=true")
	}
	flagData := result.StepOutput("get")
	if flagData == nil {
		t.Fatal("expected step output from 'get' step")
	}
	if flagData["key"] != "my-feature" {
		t.Errorf("expected key=my-feature, got %v", flagData["key"])
	}
	if flagData["on"] != true {
		t.Errorf("expected on=true, got %v", flagData["on"])
	}
}

func TestIntegration_FlagCreate(t *testing.T) {
	recorder := wftest.RecordStep("step.launchdarkly_flag_create")

	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  create_flag:
    steps:
      - name: create
        type: step.launchdarkly_flag_create
        config:
          projectKey: "my-project"
          flagKey: "new-feature"
          name: "New Feature"
          kind: "boolean"
      - name: mark_done
        type: step.set
        config:
          values:
            created: true
`),
		recorder,
	)

	result := h.ExecutePipeline("create_flag", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["created"] != true {
		t.Error("expected created=true")
	}
	calls := recorder.Calls()
	if len(calls) != 1 {
		t.Errorf("expected 1 call to launchdarkly_flag_create, got %d", len(calls))
	}
	if calls[0].Config["flagKey"] != "new-feature" {
		t.Errorf("expected flagKey=new-feature in config, got %v", calls[0].Config["flagKey"])
	}
}
