# arn:aws:iot:ap-northeast-1:111111111111:policy/policy-name
output "iot_policy" {
  value = provider::arn::iot_policy("policy-name")
}
