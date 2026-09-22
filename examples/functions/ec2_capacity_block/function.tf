# arn:aws:ec2:ap-northeast-1:111111111111:capacity-block/capacity-block-id
output "ec2_capacity_block" {
  value = provider::arn::ec2_capacity_block("capacity-block-id")
}
