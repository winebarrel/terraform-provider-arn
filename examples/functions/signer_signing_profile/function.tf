# arn:aws:signer:ap-northeast-1:111111111111:/signing-profiles/profile-name
output "signer_signing_profile" {
  value = provider::arn::signer_signing_profile("profile-name")
}
