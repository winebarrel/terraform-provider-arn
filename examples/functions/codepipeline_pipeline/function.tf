# arn:aws:codepipeline:ap-northeast-1:111111111111:pipeline-name
output "codepipeline_pipeline" {
  value = provider::arn::codepipeline_pipeline("pipeline-name")
}
