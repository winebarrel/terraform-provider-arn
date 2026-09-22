# arn:aws:resource-groups:ap-northeast-1:111111111111:group/group-name
output "ec2_group" {
  value = provider::arn::ec2_group("group-name")
}
