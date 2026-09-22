# arn:aws:ec2:ap-northeast-1:111111111111:prefix-list/prefix-list-id
output "ec2_prefix_list" {
  value = provider::arn::ec2_prefix_list("prefix-list-id")
}
