# arn:aws:mpa:ap-northeast-1:111111111111:identity-source/identity-source-id
output "mpa_identity_source" {
  value = provider::arn::mpa_identity_source("identity-source-id")
}
