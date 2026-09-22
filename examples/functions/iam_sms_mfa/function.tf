# arn:aws:iam::111111111111:sms-mfa/mfa-token-id-with-path
output "iam_sms_mfa" {
  value = provider::arn::iam_sms_mfa("mfa-token-id-with-path")
}
