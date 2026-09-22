# arn:aws:mpa:ap-northeast-1:111111111111:approval-team/approval-team-id
output "payment_cryptography_approval_team" {
  value = provider::arn::payment_cryptography_approval_team("approval-team-id")
}
