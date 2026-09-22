# arn:aws:ec2:ap-northeast-1:111111111111:fleet/fleet-id
output "ec2_fleet" {
  value = provider::arn::ec2_fleet("fleet-id")
}
