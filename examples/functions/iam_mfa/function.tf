# arn:aws:iam::111111111111:mfa/mfa-token-id-with-path
output "iam_mfa" {
  value = provider::arn::iam_mfa("mfa-token-id-with-path")
}
