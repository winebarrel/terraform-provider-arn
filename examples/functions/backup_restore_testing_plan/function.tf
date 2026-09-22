# arn:aws:backup:ap-northeast-1:111111111111:restore-testing-plan:restore-testing-plan-name-restore-testing-plan-id
output "backup_restore_testing_plan" {
  value = provider::arn::backup_restore_testing_plan("restore-testing-plan-name", "restore-testing-plan-id")
}
