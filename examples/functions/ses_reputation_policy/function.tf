# arn:aws:ses:ap-northeast-1:aws:reputation-policy/reputation-policy-name
output "ses_reputation_policy" {
  value = provider::arn::ses_reputation_policy("reputation-policy-name")
}
