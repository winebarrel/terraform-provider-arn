# arn:aws:acm:ap-northeast-1:111111111111:certificate/certificate-id
output "acm_certificate" {
  value = provider::arn::acm_certificate("certificate-id")
}
