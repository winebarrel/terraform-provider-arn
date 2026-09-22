# arn:aws:transform-custom:ap-northeast-1:111111111111:analysis/analysis-id
output "transform_custom_analysis" {
  value = provider::arn::transform_custom_analysis("analysis-id")
}
