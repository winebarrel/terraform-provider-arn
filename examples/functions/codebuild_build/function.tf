# arn:aws:codebuild:ap-northeast-1:111111111111:build/build-id
output "codebuild_build" {
  value = provider::arn::codebuild_build("build-id")
}
