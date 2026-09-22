# arn:aws:ec2:ap-northeast-1:111111111111:volume/volume-id
output "ec2_volume" {
  value = provider::arn::ec2_volume("volume-id")
}
