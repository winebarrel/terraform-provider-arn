# arn:aws:ec2:ap-northeast-1:111111111111:instance-connect-endpoint/instance-connect-endpoint-id
output "ec2_instance_connect_endpoint" {
  value = provider::arn::ec2_instance_connect_endpoint("instance-connect-endpoint-id")
}
