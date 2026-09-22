# arn:aws:voiceid:ap-northeast-1:111111111111:domain/domain-id
output "voiceid_domain" {
  value = provider::arn::voiceid_domain("domain-id")
}
