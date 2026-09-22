# arn:aws:codepipeline:ap-northeast-1:111111111111:pipeline-name/stage-name/action-name
output "codepipeline_action" {
  value = provider::arn::codepipeline_action("pipeline-name", "stage-name", "action-name")
}
