# arn:aws:rds:ap-northeast-1:111111111111:pg:parameter-group-name
output "rds_pg" {
  value = provider::arn::rds_pg("parameter-group-name")
}
