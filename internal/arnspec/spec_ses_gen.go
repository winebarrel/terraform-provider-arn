// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ses
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ses/ses.json
// Functions: 20
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ses_addon_instance", Service: "ses", Resource: "addon-instance", Template: "arn:${Partition}:ses:${Region}:${Account}:addon-instance/${AddonInstanceId}"},
		{Name: "ses_addon_subscription", Service: "ses", Resource: "addon-subscription", Template: "arn:${Partition}:ses:${Region}:${Account}:addon-subscription/${AddonSubscriptionId}"},
		{Name: "ses_configuration_set", Service: "ses", Resource: "configuration-set", Template: "arn:${Partition}:ses:${Region}:${Account}:configuration-set/${ConfigurationSetName}"},
		{Name: "ses_contact_list", Service: "ses", Resource: "contact-list", Template: "arn:${Partition}:ses:${Region}:${Account}:contact-list/${ContactListName}"},
		{Name: "ses_custom_verification_email_template", Service: "ses", Resource: "custom-verification-email-template", Template: "arn:${Partition}:ses:${Region}:${Account}:custom-verification-email-template/${TemplateName}"},
		{Name: "ses_dedicated_ip_pool", Service: "ses", Resource: "dedicated-ip-pool", Template: "arn:${Partition}:ses:${Region}:${Account}:dedicated-ip-pool/${DedicatedIPPool}"},
		{Name: "ses_deliverability_test_report", Service: "ses", Resource: "deliverability-test-report", Template: "arn:${Partition}:ses:${Region}:${Account}:deliverability-test-report/${ReportId}"},
		{Name: "ses_export_job", Service: "ses", Resource: "export-job", Template: "arn:${Partition}:ses:${Region}:${Account}:export-job/${ExportJobId}"},
		{Name: "ses_identity", Service: "ses", Resource: "identity", Template: "arn:${Partition}:ses:${Region}:${Account}:identity/${IdentityName}"},
		{Name: "ses_import_job", Service: "ses", Resource: "import-job", Template: "arn:${Partition}:ses:${Region}:${Account}:import-job/${ImportJobId}"},
		{Name: "ses_mailmanager_address_list", Service: "ses", Resource: "mailmanager-address-list", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-address-list/${AddressListId}"},
		{Name: "ses_mailmanager_archive", Service: "ses", Resource: "mailmanager-archive", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-archive/${ArchiveId}"},
		{Name: "ses_mailmanager_ingress_point", Service: "ses", Resource: "mailmanager-ingress-point", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-ingress-point/${IngressPointId}"},
		{Name: "ses_mailmanager_rule_set", Service: "ses", Resource: "mailmanager-rule-set", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-rule-set/${RuleSetId}"},
		{Name: "ses_mailmanager_smtp_relay", Service: "ses", Resource: "mailmanager-smtp-relay", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-smtp-relay/${RelayId}"},
		{Name: "ses_mailmanager_traffic_policy", Service: "ses", Resource: "mailmanager-traffic-policy", Template: "arn:${Partition}:ses:${Region}:${Account}:mailmanager-traffic-policy/${TrafficPolicyId}"},
		{Name: "ses_multi_region_endpoint", Service: "ses", Resource: "multi-region-endpoint", Template: "arn:${Partition}:ses:${Region}:${Account}:multi-region-endpoint/${EndpointName}"},
		{Name: "ses_reputation_policy", Service: "ses", Resource: "reputation-policy", Template: "arn:${Partition}:ses:${Region}:aws:reputation-policy/${ReputationPolicyName}"},
		{Name: "ses_template", Service: "ses", Resource: "template", Template: "arn:${Partition}:ses:${Region}:${Account}:template/${TemplateName}"},
		{Name: "ses_tenant", Service: "ses", Resource: "tenant", Template: "arn:${Partition}:ses:${Region}:${Account}:tenant/${TenantName}/${TenantId}"},
	})
}
