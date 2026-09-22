# arn:aws:ses:ap-northeast-1:111111111111:identity/identity-name
output "ses_identity" {
  value = provider::arn::ses_identity("identity-name")
}
