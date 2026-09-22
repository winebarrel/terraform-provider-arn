# arn:aws:codepipeline:ap-northeast-1:111111111111:pipeline-name/stage-name
output "codepipeline_stage" {
  value = provider::arn::codepipeline_stage("pipeline-name", "stage-name")
}
