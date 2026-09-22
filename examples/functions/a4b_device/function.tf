# arn:aws:a4b:ap-northeast-1:111111111111:device/resource-id
output "a4b_device" {
  value = provider::arn::a4b_device("resource-id")
}
