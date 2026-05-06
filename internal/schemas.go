package internal

import (
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// ModuleSchemas implements sdk.SchemaProvider, returning the UI/contract schema
// for every module type this plugin advertises. This makes the launchdarkly.provider
// module a "strict" contract descriptor in the wfctl audit.
func (p *launchDarklyPlugin) ModuleSchemas() []sdk.ModuleSchemaData {
	return []sdk.ModuleSchemaData{
		{
			Type:        "launchdarkly.provider",
			Label:       "LaunchDarkly Provider",
			Category:    "feature-management",
			Description: "LaunchDarkly API provider for feature flag management. Provides an authenticated HTTP client for the LaunchDarkly REST API v2.",
			ConfigFields: []sdk.ConfigField{
				{
					Name:        "apiKey",
					Type:        "string",
					Description: "LaunchDarkly API access token (required)",
					Required:    true,
				},
				{
					Name:         "apiUrl",
					Type:         "string",
					Description:  "LaunchDarkly API base URL (optional, defaults to https://app.launchdarkly.com/api/v2)",
					Required:     false,
					DefaultValue: "https://app.launchdarkly.com/api/v2",
				},
			},
		},
	}
}
