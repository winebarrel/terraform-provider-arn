# arn:aws:iot:ap-northeast-1:111111111111:index/index-name
output "iot_index" {
  value = provider::arn::iot_index("index-name")
}
