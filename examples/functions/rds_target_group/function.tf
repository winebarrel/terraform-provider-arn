# arn:aws:rds:ap-northeast-1:111111111111:target-group:target-group-id
output "rds_target_group" {
  value = provider::arn::rds_target_group("target-group-id")
}
