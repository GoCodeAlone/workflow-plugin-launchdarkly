package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// pluginManifest is a minimal representation used to validate stepSchemas and downloads coverage.
type pluginManifest struct {
	Capabilities struct {
		ModuleTypes []string `json:"moduleTypes"`
		StepTypes   []string `json:"stepTypes"`
	} `json:"capabilities"`
	StepSchemas []struct {
		Type string `json:"type"`
	} `json:"stepSchemas"`
	Downloads []struct {
		OS   string `json:"os"`
		Arch string `json:"arch"`
		URL  string `json:"url"`
	} `json:"downloads"`
}

// contractsFile is a minimal representation of plugin.contracts.json.
type contractsFile struct {
	Version   string `json:"version"`
	Contracts []struct {
		Kind string `json:"kind"`
		Type string `json:"type"`
		Mode string `json:"mode"`
	} `json:"contracts"`
}

// TestContractCoverage validates that:
//  1. Every step type advertised in capabilities has a corresponding step schema entry.
//  2. There are no duplicate step schema entries.
//  3. There are no orphan step schema entries (schemas for unadvertised step types).
//  4. The plugin implements SchemaProvider (returns at least one module schema).
//  5. plugin.contracts.json exists and has a "strict" contract for every advertised
//     module and step type (required by wfctl plugin validate --strict-contracts).
//  6. plugin.json has at least one download entry (required for type: "external").
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

	// Verify at least one download entry exists (required for type: "external").
	if len(m.Downloads) == 0 {
		t.Error("plugin.json has no downloads entries — external plugin validation will fail")
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
		t.Error("plugin.json has no stepSchemas — MCP/LSP hover docs will be unavailable")
	}

	t.Logf("step schema coverage: %d/%d step types have schemas", len(schemaTypes), len(m.Capabilities.StepTypes))

	// Verify module schema is provided via SchemaProvider.
	p := &launchDarklyPlugin{}
	moduleSchemas := p.ModuleSchemas()
	if len(moduleSchemas) == 0 {
		t.Error("plugin does not implement SchemaProvider or returns no module schemas")
	}

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

	// --- plugin.contracts.json strict contract coverage check ---
	contractsPath := filepath.Join("..", "plugin.contracts.json")
	cdata, err := os.ReadFile(contractsPath)
	if err != nil {
		t.Fatalf("plugin.contracts.json not found at %s: %v (required for wfctl --strict-contracts)", contractsPath, err)
	}
	var cf contractsFile
	if err := json.Unmarshal(cdata, &cf); err != nil {
		t.Fatalf("failed to parse plugin.contracts.json: %v", err)
	}

	// Index contract descriptors by "kind\x00type".
	contractIndex := make(map[string]string, len(cf.Contracts))
	for _, c := range cf.Contracts {
		key := c.Kind + "\x00" + c.Type
		contractIndex[key] = c.Mode
	}

	// Every advertised module type must have a strict contract descriptor.
	for _, mt := range m.Capabilities.ModuleTypes {
		mode, ok := contractIndex["module\x00"+mt]
		if !ok {
			t.Errorf("module type %q has no contract descriptor in plugin.contracts.json", mt)
		} else if mode != "strict" {
			t.Errorf("module type %q contract mode is %q, want strict", mt, mode)
		}
	}

	// Every advertised step type must have a strict contract descriptor.
	for _, st := range m.Capabilities.StepTypes {
		mode, ok := contractIndex["step\x00"+st]
		if !ok {
			t.Errorf("step type %q has no contract descriptor in plugin.contracts.json", st)
		} else if mode != "strict" {
			t.Errorf("step type %q contract mode is %q, want strict", st, mode)
		}
	}

	t.Logf("plugin.contracts.json: %d contract descriptors", len(cf.Contracts))
}
