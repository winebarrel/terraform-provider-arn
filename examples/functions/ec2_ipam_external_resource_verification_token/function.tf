# arn:aws:ec2::111111111111:ipam-external-resource-verification-token/ipam-external-resource-verification-token-id
output "ec2_ipam_external_resource_verification_token" {
  value = provider::arn::ec2_ipam_external_resource_verification_token("ipam-external-resource-verification-token-id")
}
