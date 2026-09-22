# arn:aws:iot:ap-northeast-1:111111111111:custommetric/metric-name
output "iot_custommetric" {
  value = provider::arn::iot_custommetric("metric-name")
}
