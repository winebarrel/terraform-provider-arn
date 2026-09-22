# arn:aws:iot:ap-northeast-1:111111111111:certificateprovider/certificate-provider-name
output "iot_certificateprovider" {
  value = provider::arn::iot_certificateprovider("certificate-provider-name")
}
