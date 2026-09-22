# arn:aws:osis:ap-northeast-1:111111111111:pipeline/pipeline-name
output "osis_pipeline" {
  value = provider::arn::osis_pipeline("pipeline-name")
}
