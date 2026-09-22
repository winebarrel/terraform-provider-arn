# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-group/verified-access-group-id
output "ec2_verified_access_group" {
  value = provider::arn::ec2_verified_access_group("verified-access-group-id")
}
