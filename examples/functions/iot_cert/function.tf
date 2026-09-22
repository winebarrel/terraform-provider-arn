# arn:aws:iot:ap-northeast-1:111111111111:cert/certificate
output "iot_cert" {
  value = provider::arn::iot_cert("certificate")
}
