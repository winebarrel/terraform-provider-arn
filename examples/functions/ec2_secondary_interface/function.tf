# arn:aws:ec2:ap-northeast-1:111111111111:secondary-interface/secondary-interface-id
output "ec2_secondary_interface" {
  value = provider::arn::ec2_secondary_interface("secondary-interface-id")
}
