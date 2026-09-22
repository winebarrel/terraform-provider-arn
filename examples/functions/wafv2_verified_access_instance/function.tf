# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-instance/verified-access-instance-id
output "wafv2_verified_access_instance" {
  value = provider::arn::wafv2_verified_access_instance("verified-access-instance-id")
}
