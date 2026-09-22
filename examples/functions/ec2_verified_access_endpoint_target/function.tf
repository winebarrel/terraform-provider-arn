# arn:aws:ec2:ap-northeast-1:111111111111:verified-access-endpoint-target/verified-access-endpoint-target-id
output "ec2_verified_access_endpoint_target" {
  value = provider::arn::ec2_verified_access_endpoint_target("verified-access-endpoint-target-id")
}
