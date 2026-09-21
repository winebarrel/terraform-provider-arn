// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: backup
// Source: https://servicereference.us-east-1.amazonaws.com/v1/backup/backup.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "backup_backup_access_point", Service: "backup", Resource: "backupAccessPoint", Template: "arn:${Partition}:backup:${Region}:${Account}:accesspoint/${AccessPointName}"},
		{Name: "backup_backup_plan", Service: "backup", Resource: "backupPlan", Template: "arn:${Partition}:backup:${Region}:${Account}:backup-plan:${BackupPlanId}"},
		{Name: "backup_backup_vault", Service: "backup", Resource: "backupVault", Template: "arn:${Partition}:backup:${Region}:${Account}:backup-vault:${BackupVaultName}"},
		{Name: "backup_framework", Service: "backup", Resource: "framework", Template: "arn:${Partition}:backup:${Region}:${Account}:framework:${FrameworkName}-${FrameworkId}"},
		{Name: "backup_legal_hold", Service: "backup", Resource: "legalHold", Template: "arn:${Partition}:backup:${Region}:${Account}:legal-hold:${LegalHoldId}"},
		{Name: "backup_recovery_point", Service: "backup", Resource: "recoveryPoint", Template: "arn:${Partition}:${Vendor}:${Region}:*:${ResourceType}:${RecoveryPointId}"},
		{Name: "backup_report_plan", Service: "backup", Resource: "reportPlan", Template: "arn:${Partition}:backup:${Region}:${Account}:report-plan:${ReportPlanName}-${ReportPlanId}"},
		{Name: "backup_restore_testing_plan", Service: "backup", Resource: "restoreTestingPlan", Template: "arn:${Partition}:backup:${Region}:${Account}:restore-testing-plan:${RestoreTestingPlanName}-${RestoreTestingPlanId}"},
		{Name: "backup_tiering_configuration", Service: "backup", Resource: "tieringConfiguration", Template: "arn:${Partition}:backup:${Region}:${Account}:tiering-configuration:${TieringConfigurationName}-${TieringConfigurationId}"},
	})
}
