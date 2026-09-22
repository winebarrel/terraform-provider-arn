# arn:aws:securityhub:ap-northeast-1:111111111111:hubv2/hub-v2-id
output "securityhub_hubv2" {
  value = provider::arn::securityhub_hubv2("hub-v2-id")
}
