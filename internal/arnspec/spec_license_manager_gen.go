// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: license-manager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/license-manager/license-manager.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "license_manager_grant", Service: "license-manager", Resource: "grant", Template: "arn:${Partition}:license-manager::${Account}:grant:${GrantId}"},
		{Name: "license_manager_license", Service: "license-manager", Resource: "license", Template: "arn:${Partition}:license-manager::${Account}:license:${LicenseId}"},
		{Name: "license_manager_license_asset_group", Service: "license-manager", Resource: "license-asset-group", Template: "arn:${Partition}:license-manager:${Region}:${Account}:license-asset-group:${LicenseAssetGroupId}"},
		{Name: "license_manager_license_asset_ruleset", Service: "license-manager", Resource: "license-asset-ruleset", Template: "arn:${Partition}:license-manager:${Region}:${Account}:license-asset-ruleset:${LicenseAssetRulesetId}"},
		{Name: "license_manager_license_configuration", Service: "license-manager", Resource: "license-configuration", Template: "arn:${Partition}:license-manager:${Region}:${Account}:license-configuration:${LicenseConfigurationId}"},
		{Name: "license_manager_report_generator", Service: "license-manager", Resource: "report-generator", Template: "arn:${Partition}:license-manager:${Region}:${Account}:report-generator:${ReportGeneratorId}"},
	})
}
