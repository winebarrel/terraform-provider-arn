# arn:aws:acm:ap-northeast-1:111111111111:certificate/certificate-id
output "ec2_certificate" {
  value = provider::arn::ec2_certificate("certificate-id")
}
