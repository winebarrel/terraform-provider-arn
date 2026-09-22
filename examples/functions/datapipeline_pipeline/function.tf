# arn:aws:datapipeline:ap-northeast-1:111111111111:pipeline/pipeline-id
output "datapipeline_pipeline" {
  value = provider::arn::datapipeline_pipeline("pipeline-id")
}
