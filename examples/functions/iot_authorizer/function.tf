# arn:aws:iot:ap-northeast-1:111111111111:authorizer/authorizer-name
output "iot_authorizer" {
  value = provider::arn::iot_authorizer("authorizer-name")
}
