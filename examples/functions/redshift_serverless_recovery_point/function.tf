# arn:aws:redshift-serverless:ap-northeast-1:111111111111:recoverypoint/recovery-point-id
output "redshift_serverless_recovery_point" {
  value = provider::arn::redshift_serverless_recovery_point("recovery-point-id")
}
