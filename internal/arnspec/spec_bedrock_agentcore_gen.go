// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bedrock-agentcore
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bedrock-agentcore/bedrock-agentcore.json
// Functions: 33
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bedrock_agentcore_ab_test", Service: "bedrock-agentcore", Resource: "ab-test", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:ab-test/${ABTestId}"},
		{Name: "bedrock_agentcore_apikeycredentialprovider", Service: "bedrock-agentcore", Resource: "apikeycredentialprovider", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:token-vault/${TokenVaultId}/apikeycredentialprovider/${Name}"},
		{Name: "bedrock_agentcore_batch_evaluate", Service: "bedrock-agentcore", Resource: "batch-evaluate", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:batch-evaluate/${BatchEvaluationId}"},
		{Name: "bedrock_agentcore_browser", Service: "bedrock-agentcore", Resource: "browser", Template: "arn:${Partition}:bedrock-agentcore:${Region}:aws:browser/${BrowserId}"},
		{Name: "bedrock_agentcore_browser_custom", Service: "bedrock-agentcore", Resource: "browser-custom", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:browser-custom/${BrowserId}"},
		{Name: "bedrock_agentcore_browser_profile", Service: "bedrock-agentcore", Resource: "browser-profile", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:browser-profile/${BrowserProfileId}"},
		{Name: "bedrock_agentcore_capacity_provider", Service: "bedrock-agentcore", Resource: "capacity-provider", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:capacity-provider/${CapacityProviderId}"},
		{Name: "bedrock_agentcore_code_interpreter", Service: "bedrock-agentcore", Resource: "code-interpreter", Template: "arn:${Partition}:bedrock-agentcore:${Region}:aws:code-interpreter/${CodeInterpreterId}"},
		{Name: "bedrock_agentcore_code_interpreter_custom", Service: "bedrock-agentcore", Resource: "code-interpreter-custom", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:code-interpreter-custom/${CodeInterpreterId}"},
		{Name: "bedrock_agentcore_configuration_bundle", Service: "bedrock-agentcore", Resource: "configuration-bundle", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:configuration-bundle/${ConfigurationBundleId}"},
		{Name: "bedrock_agentcore_consent_portal", Service: "bedrock-agentcore", Resource: "consent-portal", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:consent-portal/${ConsentPortalId}"},
		{Name: "bedrock_agentcore_dataset", Service: "bedrock-agentcore", Resource: "dataset", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:dataset/${DatasetId}"},
		{Name: "bedrock_agentcore_evaluator", Service: "bedrock-agentcore", Resource: "evaluator", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:evaluator/${EvaluatorId}"},
		{Name: "bedrock_agentcore_gateway", Service: "bedrock-agentcore", Resource: "gateway", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "bedrock_agentcore_harness", Service: "bedrock-agentcore", Resource: "harness", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:harness/${HarnessId}"},
		{Name: "bedrock_agentcore_harness_endpoint", Service: "bedrock-agentcore", Resource: "harness-endpoint", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:harness/${HarnessId}/harness-endpoint/${Name}"},
		{Name: "bedrock_agentcore_memory", Service: "bedrock-agentcore", Resource: "memory", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:memory/${MemoryId}"},
		{Name: "bedrock_agentcore_oauth2credentialprovider", Service: "bedrock-agentcore", Resource: "oauth2credentialprovider", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:token-vault/${TokenVaultId}/oauth2credentialprovider/${Name}"},
		{Name: "bedrock_agentcore_online_evaluation_config", Service: "bedrock-agentcore", Resource: "online-evaluation-config", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:online-evaluation-config/${OnlineEvaluationConfigId}"},
		{Name: "bedrock_agentcore_payment_manager", Service: "bedrock-agentcore", Resource: "payment-manager", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:payment-manager/${PaymentManagerId}"},
		{Name: "bedrock_agentcore_paymentcredentialprovider", Service: "bedrock-agentcore", Resource: "paymentcredentialprovider", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:token-vault/${TokenVaultId}/paymentcredentialprovider/${Name}"},
		{Name: "bedrock_agentcore_policy", Service: "bedrock-agentcore", Resource: "policy", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:policy-engine/${PolicyEngineId}/policy/${PolicyId}"},
		{Name: "bedrock_agentcore_policy_engine", Service: "bedrock-agentcore", Resource: "policy-engine", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:policy-engine/${PolicyEngineId}"},
		{Name: "bedrock_agentcore_policy_generation", Service: "bedrock-agentcore", Resource: "policy-generation", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:policy-engine/${PolicyEngineId}/policy-generation/${PolicyGenerationId}"},
		{Name: "bedrock_agentcore_recommendation", Service: "bedrock-agentcore", Resource: "recommendation", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:recommendation/${RecommendationId}"},
		{Name: "bedrock_agentcore_registry", Service: "bedrock-agentcore", Resource: "registry", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:registry/${RegistryId}"},
		{Name: "bedrock_agentcore_registry_record", Service: "bedrock-agentcore", Resource: "registry-record", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:registry/${RegistryId}/record/${RecordId}"},
		{Name: "bedrock_agentcore_runtime", Service: "bedrock-agentcore", Resource: "runtime", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:runtime/${RuntimeId}"},
		{Name: "bedrock_agentcore_runtime_endpoint", Service: "bedrock-agentcore", Resource: "runtime-endpoint", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:runtime/${RuntimeId}/runtime-endpoint/${Name}"},
		{Name: "bedrock_agentcore_token_vault", Service: "bedrock-agentcore", Resource: "token-vault", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:token-vault/${TokenVaultId}"},
		{Name: "bedrock_agentcore_web_search", Service: "bedrock-agentcore", Resource: "web-search", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:tool/web-search.v1"},
		{Name: "bedrock_agentcore_workload_identity", Service: "bedrock-agentcore", Resource: "workload-identity", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:workload-identity-directory/${DirectoryId}/workload-identity/${WorkloadIdentityName}"},
		{Name: "bedrock_agentcore_workload_identity_directory", Service: "bedrock-agentcore", Resource: "workload-identity-directory", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:workload-identity-directory/${DirectoryId}"},
	})
}
