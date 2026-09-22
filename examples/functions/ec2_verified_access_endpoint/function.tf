# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-endpoint/verified-access-endpoint-id
output "ec2_verified_access_endpoint" {
  value = provider::arn::ec2_verified_access_endpoint("verified-access-endpoint-id")
}
