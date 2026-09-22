# arn:aws:glue:ap-northeast-1:111111111111:mlTransform/transform-id
output "glue_ml_transform" {
  value = provider::arn::glue_ml_transform("transform-id")
}
