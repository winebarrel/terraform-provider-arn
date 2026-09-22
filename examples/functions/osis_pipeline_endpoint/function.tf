# arn:aws:osis:ap-northeast-1:111111111111:endpoint/endpoint-id
output "osis_pipeline_endpoint" {
  value = provider::arn::osis_pipeline_endpoint("endpoint-id")
}
