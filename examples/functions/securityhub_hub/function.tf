# arn:aws:securityhub:ap-northeast-1:111111111111:hub/default
output "securityhub_hub" {
  value = provider::arn::securityhub_hub()
}
