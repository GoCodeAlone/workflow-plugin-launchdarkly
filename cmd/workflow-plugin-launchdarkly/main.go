package main

import (
	"github.com/GoCodeAlone/workflow-plugin-launchdarkly/internal"
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

var version = "dev"

func main() {
	sdk.Serve(internal.NewLaunchDarklyPlugin(), sdk.WithBuildVersion(sdk.ResolveBuildVersion(internal.Version)))
}
