# arn:aws:iot:ap-northeast-1:111111111111:cert/certificate
output "iotwireless_cert" {
  value = provider::arn::iotwireless_cert("certificate")
}
