package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// pluginManifest is a minimal representation used to validate stepSchemas coverage.
type pluginManifest struct {
	Capabilities struct {
		StepTypes []string `json:"stepTypes"`
	} `json:"capabilities"`
	StepSchemas []struct {
		Type string `json:"type"`
	} `json:"stepSchemas"`
}

// TestContractCoverage validates that:
//  1. Every step type advertised in capabilities has a corresponding step schema entry.
//  2. There are no duplicate step schema entries.
//  3. There are no orphan step schema entries (schemas for unadvertised step types).
//  4. The plugin implements SchemaProvider (returns at least one module schema).
func TestContractCoverage(t *testing.T) {
	// Load and parse plugin.json from the repo root.
	pluginJSONPath := filepath.Join("..", "plugin.json")
	data, err := os.ReadFile(pluginJSONPath)
	if err != nil {
		t.Fatalf("plugin.json not found at %s: %v", pluginJSONPath, err)
	}
	var m pluginManifest
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to parse plugin.json: %v", err)
	}

	// Index schemas by type, detecting duplicates.
	schemaTypes := make(map[string]int, len(m.StepSchemas))
	for _, s := range m.StepSchemas {
		schemaTypes[s.Type]++
		if schemaTypes[s.Type] > 1 {
			t.Errorf("duplicate step schema entry for type %q", s.Type)
		}
	}

	// Index capabilities step types for orphan detection.
	capabilityTypes := make(map[string]bool, len(m.Capabilities.StepTypes))
	for _, st := range m.Capabilities.StepTypes {
		capabilityTypes[st] = true
	}

	// Every advertised step type must have a schema.
	for _, st := range m.Capabilities.StepTypes {
		if schemaTypes[st] == 0 {
			t.Errorf("step type %q is advertised in capabilities but has no step schema", st)
		}
	}

	// Every step schema must correspond to an advertised step type (no orphans).
	for _, s := range m.StepSchemas {
		if !capabilityTypes[s.Type] {
			t.Errorf("step schema for %q has no matching entry in capabilities.stepTypes (orphan schema)", s.Type)
		}
	}

	if len(m.StepSchemas) == 0 {
		t.Error("plugin.json has no stepSchemas — strict contract audit will report missing_step_contract_descriptor")
	}

	t.Logf("contract coverage: %d/%d step types have schemas", len(schemaTypes), len(m.Capabilities.StepTypes))

	// Verify module schema is provided via SchemaProvider.
	p := &launchDarklyPlugin{}
	moduleSchemas := p.ModuleSchemas()
	if len(moduleSchemas) == 0 {
		t.Error("plugin does not implement SchemaProvider or returns no module schemas — strict contract audit will report missing_module_contract_descriptor")
	}

	// Every module type listed in capabilities should have a schema.
	// (The plugin.json capabilities field uses legacy format for module types.)
	// For now, just ensure launchdarkly.provider is covered.
	found := false
	for _, ms := range moduleSchemas {
		if ms.Type == "launchdarkly.provider" {
			found = true
			if ms.Description == "" {
				t.Error("module schema for launchdarkly.provider has no description")
			}
			if len(ms.ConfigFields) == 0 {
				t.Error("module schema for launchdarkly.provider has no config fields")
			}
		}
	}
	if !found {
		t.Error("module schema for launchdarkly.provider not found")
	}

	t.Logf("module schemas: %d", len(moduleSchemas))
}
