# arn:aws:ec2:ap-northeast-1:111111111111:network-acl/nacl-id
output "ec2_network_acl" {
  value = provider::arn::ec2_network_acl("nacl-id")
}
