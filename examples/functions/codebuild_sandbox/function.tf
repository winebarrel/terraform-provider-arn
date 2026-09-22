# arn:aws:codebuild:ap-northeast-1:111111111111:sandbox/sandbox-id
output "codebuild_sandbox" {
  value = provider::arn::codebuild_sandbox("sandbox-id")
}
