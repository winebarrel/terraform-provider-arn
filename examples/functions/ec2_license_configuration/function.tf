# arn:aws:license-manager:ap-northeast-1:111111111111:license-configuration:license-configuration-id
output "ec2_license_configuration" {
  value = provider::arn::ec2_license_configuration("license-configuration-id")
}
