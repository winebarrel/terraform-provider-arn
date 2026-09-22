# arn:aws:codepipeline:ap-northeast-1:111111111111:webhook:webhook-name
output "codepipeline_webhook" {
  value = provider::arn::codepipeline_webhook("webhook-name")
}
