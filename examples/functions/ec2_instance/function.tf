# arn:aws:ec2:ap-northeast-1:111111111111:instance/instance-id
output "ec2_instance" {
  value = provider::arn::ec2_instance("instance-id")
}
