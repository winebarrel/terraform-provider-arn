# arn:aws:securityagent:ap-northeast-1:111111111111:application/application-id
output "securityagent_application" {
  value = provider::arn::securityagent_application("application-id")
}
