# arn:aws:iot:ap-northeast-1:111111111111:cacert/ca-certificate
output "iot_cacert" {
  value = provider::arn::iot_cacert("ca-certificate")
}
