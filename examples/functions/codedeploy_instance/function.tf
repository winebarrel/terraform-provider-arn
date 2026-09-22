# arn:aws:codedeploy:ap-northeast-1:111111111111:instance:instance-name
output "codedeploy_instance" {
  value = provider::arn::codedeploy_instance("instance-name")
}
