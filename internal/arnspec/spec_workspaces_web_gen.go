// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workspaces-web
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workspaces-web/workspaces-web.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workspaces_web_browser_settings", Service: "workspaces-web", Resource: "browserSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:browserSettings/${BrowserSettingsId}"},
		{Name: "workspaces_web_data_protection_settings", Service: "workspaces-web", Resource: "dataProtectionSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:dataProtectionSettings/${DataProtectionSettingsId}"},
		{Name: "workspaces_web_identity_provider", Service: "workspaces-web", Resource: "identityProvider", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:identityProvider/${PortalId}/${IdentityProviderId}"},
		{Name: "workspaces_web_ip_access_settings", Service: "workspaces-web", Resource: "ipAccessSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:ipAccessSettings/${IpAccessSettingsId}"},
		{Name: "workspaces_web_network_settings", Service: "workspaces-web", Resource: "networkSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:networkSettings/${NetworkSettingsId}"},
		{Name: "workspaces_web_portal", Service: "workspaces-web", Resource: "portal", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:portal/${PortalId}"},
		{Name: "workspaces_web_session_logger", Service: "workspaces-web", Resource: "sessionLogger", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:sessionLogger/${SessionLoggerId}"},
		{Name: "workspaces_web_trust_store", Service: "workspaces-web", Resource: "trustStore", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:trustStore/${TrustStoreId}"},
		{Name: "workspaces_web_user_access_logging_settings", Service: "workspaces-web", Resource: "userAccessLoggingSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:userAccessLoggingSettings/${UserAccessLoggingSettingsId}"},
		{Name: "workspaces_web_user_settings", Service: "workspaces-web", Resource: "userSettings", Template: "arn:${Partition}:workspaces-web:${Region}:${Account}:userSettings/${UserSettingsId}"},
	})
}
