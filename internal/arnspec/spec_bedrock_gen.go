// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bedrock
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bedrock/bedrock.json
// Functions: 40
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bedrock_advanced_prompt_optimization_job", Service: "bedrock", Resource: "advanced-prompt-optimization-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:advanced-prompt-optimization-job/${ResourceId}"},
		{Name: "bedrock_agent", Service: "bedrock", Resource: "agent", Template: "arn:${Partition}:bedrock:${Region}:${Account}:agent/${AgentId}"},
		{Name: "bedrock_agent_alias", Service: "bedrock", Resource: "agent-alias", Template: "arn:${Partition}:bedrock:${Region}:${Account}:agent-alias/${AgentId}/${AgentAliasId}"},
		{Name: "bedrock_application_inference_profile", Service: "bedrock", Resource: "application-inference-profile", Template: "arn:${Partition}:bedrock:${Region}:${Account}:application-inference-profile/${ResourceId}"},
		{Name: "bedrock_async_invoke", Service: "bedrock", Resource: "async-invoke", Template: "arn:${Partition}:bedrock:${Region}:${Account}:async-invoke/${ResourceId}"},
		{Name: "bedrock_automated_reasoning_policy", Service: "bedrock", Resource: "automated-reasoning-policy", Template: "arn:${Partition}:bedrock:${Region}:${Account}:automated-reasoning-policy/${AutomatedReasoningPolicyId}"},
		{Name: "bedrock_automated_reasoning_policy_version", Service: "bedrock", Resource: "automated-reasoning-policy-version", Template: "arn:${Partition}:bedrock:${Region}:${Account}:automated-reasoning-policy/${AutomatedReasoningPolicyId}:${AutomatedReasoningPolicyVersion}"},
		{Name: "bedrock_bedrock_marketplace_model_endpoint", Service: "bedrock", Resource: "bedrock-marketplace-model-endpoint", Template: "arn:${Partition}:bedrock:${Region}:${Account}:marketplace/model-endpoint/all-access"},
		{Name: "bedrock_blueprint", Service: "bedrock", Resource: "blueprint", Template: "arn:${Partition}:bedrock:${Region}:${Account}:blueprint/${BlueprintId}"},
		{Name: "bedrock_blueprint_optimization_invocation", Service: "bedrock", Resource: "blueprint-optimization-invocation", Template: "arn:${Partition}:bedrock:${Region}:${Account}:blueprint-optimization-invocation/${ResourceId}"},
		{Name: "bedrock_custom_model", Service: "bedrock", Resource: "custom-model", Template: "arn:${Partition}:bedrock:${Region}:${Account}:custom-model/${ResourceId}"},
		{Name: "bedrock_custom_model_deployment", Service: "bedrock", Resource: "custom-model-deployment", Template: "arn:${Partition}:bedrock:${Region}:${Account}:custom-model-deployment/${ResourceId}"},
		{Name: "bedrock_data_automation_invocation_job", Service: "bedrock", Resource: "data-automation-invocation-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:data-automation-invocation/${JobId}"},
		{Name: "bedrock_data_automation_library", Service: "bedrock", Resource: "data-automation-library", Template: "arn:${Partition}:bedrock:${Region}:${Account}:data-automation-library/${DataAutomationLibraryId}"},
		{Name: "bedrock_data_automation_library_ingestion_job", Service: "bedrock", Resource: "data-automation-library-ingestion-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:data-automation-library-ingestion-job/${IngestionJobId}"},
		{Name: "bedrock_data_automation_profile", Service: "bedrock", Resource: "data-automation-profile", Template: "arn:${Partition}:bedrock:${Region}:${Account}:data-automation-profile/${ProfileId}"},
		{Name: "bedrock_data_automation_project", Service: "bedrock", Resource: "data-automation-project", Template: "arn:${Partition}:bedrock:${Region}:${Account}:data-automation-project/${ProjectId}"},
		{Name: "bedrock_default_prompt_router", Service: "bedrock", Resource: "default-prompt-router", Template: "arn:${Partition}:bedrock:${Region}:${Account}:default-prompt-router/${ResourceId}"},
		{Name: "bedrock_evaluation_job", Service: "bedrock", Resource: "evaluation-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:evaluation-job/${ResourceId}"},
		{Name: "bedrock_flow", Service: "bedrock", Resource: "flow", Template: "arn:${Partition}:bedrock:${Region}:${Account}:flow/${FlowId}"},
		{Name: "bedrock_flow_alias", Service: "bedrock", Resource: "flow-alias", Template: "arn:${Partition}:bedrock:${Region}:${Account}:flow/${FlowId}/alias/${FlowAliasId}"},
		{Name: "bedrock_flow_execution", Service: "bedrock", Resource: "flow-execution", Template: "arn:${Partition}:bedrock:${Region}:${Account}:flow/${FlowId}/alias/${FlowAliasId}/execution/${FlowExecutionId}"},
		{Name: "bedrock_foundation_model", Service: "bedrock", Resource: "foundation-model", Template: "arn:${Partition}:bedrock:${Region}::foundation-model/${ResourceId}"},
		{Name: "bedrock_guardrail", Service: "bedrock", Resource: "guardrail", Template: "arn:${Partition}:bedrock:${Region}:${Account}:guardrail/${GuardrailId}"},
		{Name: "bedrock_guardrail_profile", Service: "bedrock", Resource: "guardrail-profile", Template: "arn:${Partition}:bedrock:${Region}:${Account}:guardrail-profile/${ResourceId}"},
		{Name: "bedrock_imported_model", Service: "bedrock", Resource: "imported-model", Template: "arn:${Partition}:bedrock:${Region}:${Account}:imported-model/${ResourceId}"},
		{Name: "bedrock_inference_profile", Service: "bedrock", Resource: "inference-profile", Template: "arn:${Partition}:bedrock:${Region}:${Account}:inference-profile/${ResourceId}"},
		{Name: "bedrock_knowledge_base", Service: "bedrock", Resource: "knowledge-base", Template: "arn:${Partition}:bedrock:${Region}:${Account}:knowledge-base/${KnowledgeBaseId}"},
		{Name: "bedrock_model_copy_job", Service: "bedrock", Resource: "model-copy-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:model-copy-job/${ResourceId}"},
		{Name: "bedrock_model_customization_job", Service: "bedrock", Resource: "model-customization-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:model-customization-job/${ResourceId}"},
		{Name: "bedrock_model_evaluation_job", Service: "bedrock", Resource: "model-evaluation-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:model-evaluation-job/${ResourceId}"},
		{Name: "bedrock_model_import_job", Service: "bedrock", Resource: "model-import-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:model-import-job/${ResourceId}"},
		{Name: "bedrock_model_invocation_job", Service: "bedrock", Resource: "model-invocation-job", Template: "arn:${Partition}:bedrock:${Region}:${Account}:model-invocation-job/${JobIdentifier}"},
		{Name: "bedrock_project", Service: "bedrock", Resource: "project", Template: "arn:${Partition}:bedrock:${Region}:${Account}:project/${ResourceId}"},
		{Name: "bedrock_prompt", Service: "bedrock", Resource: "prompt", Template: "arn:${Partition}:bedrock:${Region}:${Account}:prompt/${PromptId}"},
		{Name: "bedrock_prompt_router", Service: "bedrock", Resource: "prompt-router", Template: "arn:${Partition}:bedrock:${Region}:${Account}:prompt-router/${ResourceId}"},
		{Name: "bedrock_prompt_version", Service: "bedrock", Resource: "prompt-version", Template: "arn:${Partition}:bedrock:${Region}:${Account}:prompt/${PromptId}:${PromptVersion}"},
		{Name: "bedrock_provisioned_model", Service: "bedrock", Resource: "provisioned-model", Template: "arn:${Partition}:bedrock:${Region}:${Account}:provisioned-model/${ResourceId}"},
		{Name: "bedrock_session", Service: "bedrock", Resource: "session", Template: "arn:${Partition}:bedrock:${Region}:${Account}:session/${SessionId}"},
		{Name: "bedrock_system_tool", Service: "bedrock", Resource: "system-tool", Template: "arn:${Partition}:bedrock::${Account}:system-tool/${ResourceId}"},
	})
}
