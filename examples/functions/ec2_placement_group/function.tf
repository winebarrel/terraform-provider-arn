# arn:aws:ec2:ap-northeast-1:111111111111:placement-group/placement-group-name
output "ec2_placement_group" {
  value = provider::arn::ec2_placement_group("placement-group-name")
}
