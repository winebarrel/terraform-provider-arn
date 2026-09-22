# arn:aws:ec2:ap-northeast-1:111111111111:launch-template/launch-template-id
output "ec2_launch_template" {
  value = provider::arn::ec2_launch_template("launch-template-id")
}
