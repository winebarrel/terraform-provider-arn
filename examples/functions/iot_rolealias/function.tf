# arn:aws:iot:ap-northeast-1:111111111111:rolealias/role-alias
output "iot_rolealias" {
  value = provider::arn::iot_rolealias("role-alias")
}
