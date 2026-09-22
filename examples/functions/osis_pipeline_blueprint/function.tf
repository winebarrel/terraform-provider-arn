# arn:aws:osis:ap-northeast-1:111111111111:blueprint/blueprint-name
output "osis_pipeline_blueprint" {
  value = provider::arn::osis_pipeline_blueprint("blueprint-name")
}
