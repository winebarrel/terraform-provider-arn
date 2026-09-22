# arn:aws:iotmanagedintegrations:ap-northeast-1:111111111111:credential-locker/identifier
output "iotmanagedintegrations_credential_locker" {
  value = provider::arn::iotmanagedintegrations_credential_locker("identifier")
}
