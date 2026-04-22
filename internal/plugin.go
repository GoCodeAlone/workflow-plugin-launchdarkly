// Package internal implements the workflow-plugin-launchdarkly plugin.
package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// Version is set at build time via -ldflags
// "-X github.com/GoCodeAlone/workflow-plugin-launchdarkly/internal.Version=X.Y.Z".
// Default is a bare semver so plugin loaders that validate semver accept
// unreleased dev builds; goreleaser overrides with the real release tag.
var Version = "0.0.0"

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
		Version:     Version,
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
