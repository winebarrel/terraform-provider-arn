# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/integrations/uri
output "profile_integrations" {
  value = provider::arn::profile_integrations("domain-name", "uri")
}
