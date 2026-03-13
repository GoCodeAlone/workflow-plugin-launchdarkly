// Package internal implements the workflow-plugin-launchdarkly plugin.
package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// launchDarklyPlugin implements sdk.PluginProvider, sdk.ModuleProvider, and sdk.StepProvider.
type launchDarklyPlugin struct{}

// NewLaunchDarklyPlugin returns a new launchDarklyPlugin instance.
func NewLaunchDarklyPlugin() sdk.PluginProvider {
	return &launchDarklyPlugin{}
}

// Manifest returns plugin metadata.
func (p *launchDarklyPlugin) Manifest() sdk.PluginManifest {
	return sdk.PluginManifest{
		Name:        "workflow-plugin-launchdarkly",
		Version:     "0.1.0",
		Author:      "GoCodeAlone",
		Description: "LaunchDarkly feature flag management plugin (~100 step types across all LaunchDarkly APIs)",
	}
}

// ModuleTypes returns the module type names this plugin provides.
func (p *launchDarklyPlugin) ModuleTypes() []string {
	return []string{"launchdarkly.provider"}
}

// CreateModule creates a module instance of the given type.
func (p *launchDarklyPlugin) CreateModule(typeName, name string, config map[string]any) (sdk.ModuleInstance, error) {
	switch typeName {
	case "launchdarkly.provider":
		m, err := newLaunchDarklyModule(name, config)
		if err != nil {
			return nil, err
		}
		return m, nil
	default:
		return nil, fmt.Errorf("launchdarkly plugin: unknown module type %q", typeName)
	}
}

// StepTypes returns the step type names this plugin provides.
func (p *launchDarklyPlugin) StepTypes() []string {
	return allStepTypes()
}

// CreateStep creates a step instance of the given type.
func (p *launchDarklyPlugin) CreateStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	return createStep(typeName, name, config)
}
