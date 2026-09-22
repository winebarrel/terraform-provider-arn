# arn:aws:codepipeline:ap-northeast-1:111111111111:actiontype:owner/category/provider/version
output "codepipeline_actiontype" {
  value = provider::arn::codepipeline_actiontype("owner", "category", "provider", "version")
}
