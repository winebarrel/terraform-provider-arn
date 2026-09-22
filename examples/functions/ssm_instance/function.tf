# arn:aws:ec2:ap-northeast-1:111111111111:instance/instance-id
output "ssm_instance" {
  value = provider::arn::ssm_instance("instance-id")
}
