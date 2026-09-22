# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/experience/experience-id
output "kendra_experience" {
  value = provider::arn::kendra_experience("index-id", "experience-id")
}
