# arn:aws:codebuild:ap-northeast-1:111111111111:fleet/fleet-name:fleet-id
output "codebuild_fleet" {
  value = provider::arn::codebuild_fleet("fleet-name", "fleet-id")
}
