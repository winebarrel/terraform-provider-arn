# arn:aws:ec2:ap-northeast-1:111111111111:key-pair/key-pair-name
output "ec2_key_pair" {
  value = provider::arn::ec2_key_pair("key-pair-name")
}
