# arn:aws:ec2:ap-northeast-1:111111111111:dedicated-host/dedicated-host-id
output "ec2_dedicated_host" {
  value = provider::arn::ec2_dedicated_host("dedicated-host-id")
}
