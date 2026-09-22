# arn:aws:transform-custom:ap-northeast-1:111111111111:remediation/remediation-id
output "transform_custom_remediation" {
  value = provider::arn::transform_custom_remediation("remediation-id")
}
