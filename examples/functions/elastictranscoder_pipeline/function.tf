# arn:aws:elastictranscoder:ap-northeast-1:111111111111:pipeline/pipeline-id
output "elastictranscoder_pipeline" {
  value = provider::arn::elastictranscoder_pipeline("pipeline-id")
}
