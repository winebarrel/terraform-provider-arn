# arn:aws:ec2:ap-northeast-1:111111111111:spot-instances-request/spot-instance-request-id
output "ec2_spot_instances_request" {
  value = provider::arn::ec2_spot_instances_request("spot-instance-request-id")
}
