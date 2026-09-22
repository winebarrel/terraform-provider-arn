# arn:aws:securityagent:ap-northeast-1:111111111111:target-domain/target-domain-id
output "securityagent_target_domain" {
  value = provider::arn::securityagent_target_domain("target-domain-id")
}
