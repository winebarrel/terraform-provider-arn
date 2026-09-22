# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/access-control-configuration/access-control-configuration-id
output "kendra_access_control_configuration" {
  value = provider::arn::kendra_access_control_configuration("index-id", "access-control-configuration-id")
}
