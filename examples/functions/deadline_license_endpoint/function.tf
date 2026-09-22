# arn:aws:deadline:ap-northeast-1:111111111111:license-endpoint/license-endpoint-id
output "deadline_license_endpoint" {
  value = provider::arn::deadline_license_endpoint("license-endpoint-id")
}
