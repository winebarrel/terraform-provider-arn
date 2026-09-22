# arn:aws:mediaconvert:ap-northeast-1:111111111111:certificates/certificate-arn
output "mediaconvert_certificate_association" {
  value = provider::arn::mediaconvert_certificate_association("certificate-arn")
}
