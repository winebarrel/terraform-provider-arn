// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: quicksight
// Source: https://servicereference.us-east-1.amazonaws.com/v1/quicksight/quicksight.json
// Functions: 38
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "quicksight_account", Service: "quicksight", Resource: "account", Template: "arn:${Partition}:quicksight:${Region}:${Account}:account/${ResourceId}"},
		{Name: "quicksight_actionconnector", Service: "quicksight", Resource: "actionconnector", Template: "arn:${Partition}:quicksight:${Region}:${Account}:action-connector/${ResourceId}"},
		{Name: "quicksight_agent", Service: "quicksight", Resource: "agent", Template: "arn:${Partition}:quicksight:${Region}:${Account}:agent/${ResourceId}"},
		{Name: "quicksight_analysis", Service: "quicksight", Resource: "analysis", Template: "arn:${Partition}:quicksight:${Region}:${Account}:analysis/${ResourceId}"},
		{Name: "quicksight_app", Service: "quicksight", Resource: "app", Template: "arn:${Partition}:quicksight:${Region}:${Account}:app/${ResourceId}"},
		{Name: "quicksight_approval_policy", Service: "quicksight", Resource: "approvalPolicy", Template: "arn:${Partition}:quicksight:${Region}:${Account}:approval-policy/${ResourceId}"},
		{Name: "quicksight_asset_bundle_export_job", Service: "quicksight", Resource: "assetBundleExportJob", Template: "arn:${Partition}:quicksight:${Region}:${Account}:asset-bundle-export-job/${ResourceId}"},
		{Name: "quicksight_asset_bundle_import_job", Service: "quicksight", Resource: "assetBundleImportJob", Template: "arn:${Partition}:quicksight:${Region}:${Account}:asset-bundle-import-job/${ResourceId}"},
		{Name: "quicksight_assignment", Service: "quicksight", Resource: "assignment", Template: "arn:${Partition}:quicksight::${Account}:assignment/${ResourceId}"},
		{Name: "quicksight_automation", Service: "quicksight", Resource: "automation", Template: "arn:${Partition}:quicksight:${Region}:${Account}:automation-group/${AutomationGroupId}/automation/${ResourceId}"},
		{Name: "quicksight_automation_group", Service: "quicksight", Resource: "automationGroup", Template: "arn:${Partition}:quicksight:${Region}:${Account}:automation-group/${ResourceId}"},
		{Name: "quicksight_automation_job", Service: "quicksight", Resource: "automationJob", Template: "arn:${Partition}:quicksight:${Region}:${Account}:automation-group/${AutomationGroupId}/automation/${AutomationId}/job/${ResourceId}"},
		{Name: "quicksight_brand", Service: "quicksight", Resource: "brand", Template: "arn:${Partition}:quicksight:${Region}:${Account}:brand/${ResourceId}"},
		{Name: "quicksight_customization", Service: "quicksight", Resource: "customization", Template: "arn:${Partition}:quicksight:${Region}:${Account}:customization/${ResourceId}"},
		{Name: "quicksight_custompermissions", Service: "quicksight", Resource: "custompermissions", Template: "arn:${Partition}:quicksight:${Region}:${Account}:custompermissions/${ResourceId}"},
		{Name: "quicksight_dashboard", Service: "quicksight", Resource: "dashboard", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dashboard/${ResourceId}"},
		{Name: "quicksight_dashboard_snapshot_job", Service: "quicksight", Resource: "dashboardSnapshotJob", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dashboard/${DashboardId}/snapshot-job/${ResourceId}"},
		{Name: "quicksight_dataset", Service: "quicksight", Resource: "dataset", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dataset/${ResourceId}"},
		{Name: "quicksight_datasource", Service: "quicksight", Resource: "datasource", Template: "arn:${Partition}:quicksight:${Region}:${Account}:datasource/${ResourceId}"},
		{Name: "quicksight_dlp_setting", Service: "quicksight", Resource: "dlpSetting", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dlpsetting/${ResourceId}"},
		{Name: "quicksight_email_customization_template", Service: "quicksight", Resource: "emailCustomizationTemplate", Template: "arn:${Partition}:quicksight:${Region}:${Account}:email-customization-template/${ResourceId}"},
		{Name: "quicksight_extension", Service: "quicksight", Resource: "extension", Template: "arn:${Partition}:quicksight:${Region}:${Account}:extension/${ResourceId}"},
		{Name: "quicksight_extensionaccess", Service: "quicksight", Resource: "extensionaccess", Template: "arn:${Partition}:quicksight:${Region}:${Account}:extension-access/${ResourceId}"},
		{Name: "quicksight_flow", Service: "quicksight", Resource: "flow", Template: "arn:${Partition}:quicksight:${Region}:${Account}:flow/${ResourceId}"},
		{Name: "quicksight_folder", Service: "quicksight", Resource: "folder", Template: "arn:${Partition}:quicksight:${Region}:${Account}:folder/${ResourceId}"},
		{Name: "quicksight_group", Service: "quicksight", Resource: "group", Template: "arn:${Partition}:quicksight:${Region}:${Account}:group/${ResourceId}"},
		{Name: "quicksight_ingestion", Service: "quicksight", Resource: "ingestion", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dataset/${DatasetId}/ingestion/${ResourceId}"},
		{Name: "quicksight_knowledge_base", Service: "quicksight", Resource: "knowledgeBase", Template: "arn:${Partition}:quicksight:${Region}:${Account}:knowledge-base/${ResourceId}"},
		{Name: "quicksight_limits_profile", Service: "quicksight", Resource: "limitsProfile", Template: "arn:${Partition}:quicksight:${Region}:${Account}:limits-profile/${ResourceId}"},
		{Name: "quicksight_namespace", Service: "quicksight", Resource: "namespace", Template: "arn:${Partition}:quicksight:${Region}:${Account}:namespace/${ResourceId}"},
		{Name: "quicksight_oauth_client_application", Service: "quicksight", Resource: "oauthClientApplication", Template: "arn:${Partition}:quicksight:${Region}:${Account}:oauthClientApplication/${ResourceId}"},
		{Name: "quicksight_refreshschedule", Service: "quicksight", Resource: "refreshschedule", Template: "arn:${Partition}:quicksight:${Region}:${Account}:dataset/${DatasetId}/refresh-schedule/${ResourceId}"},
		{Name: "quicksight_space", Service: "quicksight", Resource: "space", Template: "arn:${Partition}:quicksight:${Region}:${Account}:space/${ResourceId}"},
		{Name: "quicksight_template", Service: "quicksight", Resource: "template", Template: "arn:${Partition}:quicksight:${Region}:${Account}:template/${ResourceId}"},
		{Name: "quicksight_theme", Service: "quicksight", Resource: "theme", Template: "arn:${Partition}:quicksight:${Region}:${Account}:theme/${ResourceId}"},
		{Name: "quicksight_topic", Service: "quicksight", Resource: "topic", Template: "arn:${Partition}:quicksight:${Region}:${Account}:topic/${ResourceId}"},
		{Name: "quicksight_user", Service: "quicksight", Resource: "user", Template: "arn:${Partition}:quicksight:${Region}:${Account}:user/${ResourceId}"},
		{Name: "quicksight_vpcconnection", Service: "quicksight", Resource: "vpcconnection", Template: "arn:${Partition}:quicksight:${Region}:${Account}:vpcConnection/${ResourceId}"},
	})
}
