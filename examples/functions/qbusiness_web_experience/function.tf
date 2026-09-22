# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/web-experience/web-experience-id
output "qbusiness_web_experience" {
  value = provider::arn::qbusiness_web_experience("application-id", "web-experience-id")
}
