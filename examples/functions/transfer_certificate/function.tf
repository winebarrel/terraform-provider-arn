# arn:aws:transfer:ap-northeast-1:111111111111:certificate/certificate-id
output "transfer_certificate" {
  value = provider::arn::transfer_certificate("certificate-id")
}
