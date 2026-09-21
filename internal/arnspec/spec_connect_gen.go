// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: connect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/connect/connect.json
// Functions: 44
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "connect_agent_status", Service: "connect", Resource: "agent-status", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/agent-state/${AgentStatusId}"},
		{Name: "connect_ai_agent", Service: "connect", Resource: "ai-agent", Template: "arn:${Partition}:wisdom:${Region}:${Account}:ai-agent/${AssistantId}/${AIAgentId}:${Version}"},
		{Name: "connect_attached_file", Service: "connect", Resource: "attached-file", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/file/${FileId}"},
		{Name: "connect_authentication_profile", Service: "connect", Resource: "authentication-profile", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/authentication-profile/${AuthenticationProfileId}"},
		{Name: "connect_aws_managed_view", Service: "connect", Resource: "aws-managed-view", Template: "arn:${Partition}:connect:${Region}:aws:view/${ViewId}"},
		{Name: "connect_contact", Service: "connect", Resource: "contact", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/contact/${ContactId}"},
		{Name: "connect_contact_evaluation", Service: "connect", Resource: "contact-evaluation", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/contact-evaluation/${EvaluationId}"},
		{Name: "connect_contact_flow", Service: "connect", Resource: "contact-flow", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/contact-flow/${ContactFlowId}"},
		{Name: "connect_contact_flow_module", Service: "connect", Resource: "contact-flow-module", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/flow-module/${ContactFlowModuleId}"},
		{Name: "connect_customer_managed_view", Service: "connect", Resource: "customer-managed-view", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/view/${ViewId}"},
		{Name: "connect_customer_managed_view_version", Service: "connect", Resource: "customer-managed-view-version", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/view/${ViewId}:${ViewVersion}"},
		{Name: "connect_data_table", Service: "connect", Resource: "data-table", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/data-table/${DataTableId}"},
		{Name: "connect_email_address", Service: "connect", Resource: "email-address", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/email-address/${EmailAddressId}"},
		{Name: "connect_evaluation_form", Service: "connect", Resource: "evaluation-form", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/evaluation-form/${FormId}"},
		{Name: "connect_extraction_definition", Service: "connect", Resource: "extraction-definition", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/extraction-definition/${ExtractionDefinitionId}"},
		{Name: "connect_hierarchy_group", Service: "connect", Resource: "hierarchy-group", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/agent-group/${HierarchyGroupId}"},
		{Name: "connect_hours_of_operation", Service: "connect", Resource: "hours-of-operation", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/operating-hours/${HoursOfOperationId}"},
		{Name: "connect_instance", Service: "connect", Resource: "instance", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}"},
		{Name: "connect_integration_association", Service: "connect", Resource: "integration-association", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/integration-association/${IntegrationAssociationId}"},
		{Name: "connect_legacy_phone_number", Service: "connect", Resource: "legacy-phone-number", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/phone-number/${PhoneNumberId}"},
		{Name: "connect_metric", Service: "connect", Resource: "metric", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/metric/${MetricId}"},
		{Name: "connect_notification", Service: "connect", Resource: "notification", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/notification/${NotificationId}"},
		{Name: "connect_phone_number", Service: "connect", Resource: "phone-number", Template: "arn:${Partition}:connect:${Region}:${Account}:phone-number/${PhoneNumberId}"},
		{Name: "connect_prompt", Service: "connect", Resource: "prompt", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/prompt/${PromptId}"},
		{Name: "connect_qualified_aws_managed_view", Service: "connect", Resource: "qualified-aws-managed-view", Template: "arn:${Partition}:connect:${Region}:aws:view/${ViewId}:${ViewQualifier}"},
		{Name: "connect_qualified_customer_managed_view", Service: "connect", Resource: "qualified-customer-managed-view", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/view/${ViewId}:${ViewQualifier}"},
		{Name: "connect_qualified_metric", Service: "connect", Resource: "qualified-metric", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/metric/${MetricId}:${MetricQualifier}"},
		{Name: "connect_queue", Service: "connect", Resource: "queue", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/queue/${QueueId}"},
		{Name: "connect_quick_connect", Service: "connect", Resource: "quick-connect", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/transfer-destination/${QuickConnectId}"},
		{Name: "connect_routing_profile", Service: "connect", Resource: "routing-profile", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/routing-profile/${RoutingProfileId}"},
		{Name: "connect_rule", Service: "connect", Resource: "rule", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/rule/${RuleId}"},
		{Name: "connect_security_profile", Service: "connect", Resource: "security-profile", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/security-profile/${SecurityProfileId}"},
		{Name: "connect_task_template", Service: "connect", Resource: "task-template", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/task-template/${TaskTemplateId}"},
		{Name: "connect_traffic_distribution_group", Service: "connect", Resource: "traffic-distribution-group", Template: "arn:${Partition}:connect:${Region}:${Account}:traffic-distribution-group/${TrafficDistributionGroupId}"},
		{Name: "connect_use_case", Service: "connect", Resource: "use-case", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/use-case/${UseCaseId}"},
		{Name: "connect_user", Service: "connect", Resource: "user", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/agent/${UserId}"},
		{Name: "connect_vocabulary", Service: "connect", Resource: "vocabulary", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/vocabulary/${VocabularyId}"},
		{Name: "connect_wildcard_agent_status", Service: "connect", Resource: "wildcard-agent-status", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/agent-state/*"},
		{Name: "connect_wildcard_contact_flow", Service: "connect", Resource: "wildcard-contact-flow", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/contact-flow/*"},
		{Name: "connect_wildcard_legacy_phone_number", Service: "connect", Resource: "wildcard-legacy-phone-number", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/phone-number/*"},
		{Name: "connect_wildcard_phone_number", Service: "connect", Resource: "wildcard-phone-number", Template: "arn:${Partition}:connect:${Region}:${Account}:phone-number/*"},
		{Name: "connect_wildcard_queue", Service: "connect", Resource: "wildcard-queue", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/queue/*"},
		{Name: "connect_wildcard_quick_connect", Service: "connect", Resource: "wildcard-quick-connect", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/transfer-destination/*"},
		{Name: "connect_workspace", Service: "connect", Resource: "workspace", Template: "arn:${Partition}:connect:${Region}:${Account}:instance/${InstanceId}/workspace/${WorkspaceId}"},
	})
}
