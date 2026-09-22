# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-policy/verified-access-policy-id
output "ec2_verified_access_policy" {
  value = provider::arn::ec2_verified_access_policy("verified-access-policy-id")
}
