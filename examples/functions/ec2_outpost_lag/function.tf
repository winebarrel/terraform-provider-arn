# arn:aws:ec2:ap-northeast-1:111111111111:outpost-lag/outpost-lag-id
output "ec2_outpost_lag" {
  value = provider::arn::ec2_outpost_lag("outpost-lag-id")
}
