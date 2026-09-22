# arn:aws:ses:ap-northeast-1:111111111111:multi-region-endpoint/endpoint-name
output "ses_multi_region_endpoint" {
  value = provider::arn::ses_multi_region_endpoint("endpoint-name")
}
