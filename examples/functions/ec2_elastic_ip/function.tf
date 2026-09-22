# arn:aws:ec2:ap-northeast-1:111111111111:elastic-ip/allocation-id
output "ec2_elastic_ip" {
  value = provider::arn::ec2_elastic_ip("allocation-id")
}
