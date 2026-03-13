package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// stepConstructor is a function that creates a StepInstance.
type stepConstructor func(name string, config map[string]any) (sdk.StepInstance, error)

// stepRegistry maps step type strings to constructor functions.
var stepRegistry = map[string]stepConstructor{
	// Feature Flags
	"step.launchdarkly_flag_list":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagListStep(n, c) },
	"step.launchdarkly_flag_get":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagGetStep(n, c) },
	"step.launchdarkly_flag_create":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagCreateStep(n, c) },
	"step.launchdarkly_flag_update":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagUpdateStep(n, c) },
	"step.launchdarkly_flag_delete":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagDeleteStep(n, c) },
	"step.launchdarkly_flag_copy":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagCopyStep(n, c) },
	"step.launchdarkly_flag_status_get":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagStatusGetStep(n, c) },
	"step.launchdarkly_flag_status_list": func(n string, c map[string]any) (sdk.StepInstance, error) { return newFlagStatusListStep(n, c) },

	// Projects
	"step.launchdarkly_project_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newProjectListStep(n, c) },
	"step.launchdarkly_project_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newProjectGetStep(n, c) },
	"step.launchdarkly_project_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newProjectCreateStep(n, c) },
	"step.launchdarkly_project_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newProjectUpdateStep(n, c) },
	"step.launchdarkly_project_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newProjectDeleteStep(n, c) },

	// Environments
	"step.launchdarkly_environment_list":              func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentListStep(n, c) },
	"step.launchdarkly_environment_get":               func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentGetStep(n, c) },
	"step.launchdarkly_environment_create":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentCreateStep(n, c) },
	"step.launchdarkly_environment_update":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentUpdateStep(n, c) },
	"step.launchdarkly_environment_delete":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentDeleteStep(n, c) },
	"step.launchdarkly_environment_reset_sdk_key":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentResetSDKKeyStep(n, c) },
	"step.launchdarkly_environment_reset_mobile_key":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnvironmentResetMobileKeyStep(n, c) },

	// Segments
	"step.launchdarkly_segment_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newSegmentListStep(n, c) },
	"step.launchdarkly_segment_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newSegmentGetStep(n, c) },
	"step.launchdarkly_segment_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newSegmentCreateStep(n, c) },
	"step.launchdarkly_segment_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newSegmentUpdateStep(n, c) },
	"step.launchdarkly_segment_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newSegmentDeleteStep(n, c) },

	// Contexts
	"step.launchdarkly_context_list":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextListStep(n, c) },
	"step.launchdarkly_context_get":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextGetStep(n, c) },
	"step.launchdarkly_context_search":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextSearchStep(n, c) },
	"step.launchdarkly_context_kind_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextKindListStep(n, c) },
	"step.launchdarkly_context_kind_upsert": func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextKindUpsertStep(n, c) },
	"step.launchdarkly_context_evaluate":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newContextEvaluateStep(n, c) },

	// Metrics
	"step.launchdarkly_metric_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newMetricListStep(n, c) },
	"step.launchdarkly_metric_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newMetricGetStep(n, c) },
	"step.launchdarkly_metric_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMetricCreateStep(n, c) },
	"step.launchdarkly_metric_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMetricUpdateStep(n, c) },
	"step.launchdarkly_metric_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMetricDeleteStep(n, c) },

	// Experiments
	"step.launchdarkly_experiment_list":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newExperimentListStep(n, c) },
	"step.launchdarkly_experiment_get":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newExperimentGetStep(n, c) },
	"step.launchdarkly_experiment_create":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newExperimentCreateStep(n, c) },
	"step.launchdarkly_experiment_update":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newExperimentUpdateStep(n, c) },
	"step.launchdarkly_experiment_results_get": func(n string, c map[string]any) (sdk.StepInstance, error) { return newExperimentResultsGetStep(n, c) },

	// Approvals
	"step.launchdarkly_approval_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalListStep(n, c) },
	"step.launchdarkly_approval_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalGetStep(n, c) },
	"step.launchdarkly_approval_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalCreateStep(n, c) },
	"step.launchdarkly_approval_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalDeleteStep(n, c) },
	"step.launchdarkly_approval_apply":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalApplyStep(n, c) },
	"step.launchdarkly_approval_review": func(n string, c map[string]any) (sdk.StepInstance, error) { return newApprovalReviewStep(n, c) },

	// Scheduled Changes
	"step.launchdarkly_scheduled_change_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newScheduledChangeListStep(n, c) },
	"step.launchdarkly_scheduled_change_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newScheduledChangeCreateStep(n, c) },
	"step.launchdarkly_scheduled_change_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newScheduledChangeUpdateStep(n, c) },
	"step.launchdarkly_scheduled_change_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newScheduledChangeDeleteStep(n, c) },

	// Flag Triggers
	"step.launchdarkly_trigger_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerListStep(n, c) },
	"step.launchdarkly_trigger_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerCreateStep(n, c) },
	"step.launchdarkly_trigger_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerGetStep(n, c) },
	"step.launchdarkly_trigger_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerUpdateStep(n, c) },
	"step.launchdarkly_trigger_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerDeleteStep(n, c) },

	// Workflows
	"step.launchdarkly_workflow_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newWorkflowListStep(n, c) },
	"step.launchdarkly_workflow_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newWorkflowGetStep(n, c) },
	"step.launchdarkly_workflow_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWorkflowCreateStep(n, c) },
	"step.launchdarkly_workflow_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWorkflowDeleteStep(n, c) },

	// Audit Log
	"step.launchdarkly_audit_log_list": func(n string, c map[string]any) (sdk.StepInstance, error) { return newAuditLogListStep(n, c) },
	"step.launchdarkly_audit_log_get":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newAuditLogGetStep(n, c) },

	// Members
	"step.launchdarkly_member_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newMemberListStep(n, c) },
	"step.launchdarkly_member_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newMemberGetStep(n, c) },
	"step.launchdarkly_member_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMemberCreateStep(n, c) },
	"step.launchdarkly_member_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMemberUpdateStep(n, c) },
	"step.launchdarkly_member_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMemberDeleteStep(n, c) },

	// Teams
	"step.launchdarkly_team_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newTeamListStep(n, c) },
	"step.launchdarkly_team_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newTeamGetStep(n, c) },
	"step.launchdarkly_team_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTeamCreateStep(n, c) },
	"step.launchdarkly_team_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTeamUpdateStep(n, c) },
	"step.launchdarkly_team_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTeamDeleteStep(n, c) },

	// Custom Roles
	"step.launchdarkly_role_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleListStep(n, c) },
	"step.launchdarkly_role_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleGetStep(n, c) },
	"step.launchdarkly_role_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleCreateStep(n, c) },
	"step.launchdarkly_role_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleUpdateStep(n, c) },
	"step.launchdarkly_role_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleDeleteStep(n, c) },

	// Access Tokens
	"step.launchdarkly_token_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenListStep(n, c) },
	"step.launchdarkly_token_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenGetStep(n, c) },
	"step.launchdarkly_token_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenCreateStep(n, c) },
	"step.launchdarkly_token_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenUpdateStep(n, c) },
	"step.launchdarkly_token_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenDeleteStep(n, c) },
	"step.launchdarkly_token_reset":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newTokenResetStep(n, c) },

	// Webhooks
	"step.launchdarkly_webhook_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebhookListStep(n, c) },
	"step.launchdarkly_webhook_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebhookGetStep(n, c) },
	"step.launchdarkly_webhook_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebhookCreateStep(n, c) },
	"step.launchdarkly_webhook_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebhookUpdateStep(n, c) },
	"step.launchdarkly_webhook_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebhookDeleteStep(n, c) },

	// Relay Proxy
	"step.launchdarkly_relay_config_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newRelayConfigListStep(n, c) },
	"step.launchdarkly_relay_config_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newRelayConfigGetStep(n, c) },
	"step.launchdarkly_relay_config_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRelayConfigCreateStep(n, c) },
	"step.launchdarkly_relay_config_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRelayConfigUpdateStep(n, c) },
	"step.launchdarkly_relay_config_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRelayConfigDeleteStep(n, c) },

	// Release Pipelines
	"step.launchdarkly_release_pipeline_list":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleasePipelineListStep(n, c) },
	"step.launchdarkly_release_pipeline_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleasePipelineGetStep(n, c) },
	"step.launchdarkly_release_pipeline_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleasePipelineCreateStep(n, c) },
	"step.launchdarkly_release_pipeline_update": func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleasePipelineUpdateStep(n, c) },
	"step.launchdarkly_release_pipeline_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleasePipelineDeleteStep(n, c) },

	// Code References
	"step.launchdarkly_code_ref_repo_list":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newCodeRefRepoListStep(n, c) },
	"step.launchdarkly_code_ref_repo_create":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCodeRefRepoCreateStep(n, c) },
	"step.launchdarkly_code_ref_repo_delete":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCodeRefRepoDeleteStep(n, c) },
	"step.launchdarkly_code_ref_extinction_list": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCodeRefExtinctionListStep(n, c) },
}

// createStep dispatches to the appropriate step constructor.
func createStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	constructor, ok := stepRegistry[typeName]
	if !ok {
		return nil, fmt.Errorf("launchdarkly plugin: unknown step type %q", typeName)
	}
	return constructor(name, config)
}

// allStepTypes returns all registered step type strings.
func allStepTypes() []string {
	types := make([]string, 0, len(stepRegistry))
	for k := range stepRegistry {
		types = append(types, k)
	}
	return types
}
