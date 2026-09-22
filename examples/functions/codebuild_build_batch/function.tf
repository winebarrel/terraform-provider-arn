# arn:aws:codebuild:ap-northeast-1:111111111111:build-batch/build-batch-id
output "codebuild_build_batch" {
  value = provider::arn::codebuild_build_batch("build-batch-id")
}
