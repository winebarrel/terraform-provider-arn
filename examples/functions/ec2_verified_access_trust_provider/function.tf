# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-trust-provider/verified-access-trust-provider-id
output "ec2_verified_access_trust_provider" {
  value = provider::arn::ec2_verified_access_trust_provider("verified-access-trust-provider-id")
}
