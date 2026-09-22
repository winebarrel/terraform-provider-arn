# arn:aws:iot:ap-northeast-1:111111111111:package/package-name
output "iot_package" {
  value = provider::arn::iot_package("package-name")
}
