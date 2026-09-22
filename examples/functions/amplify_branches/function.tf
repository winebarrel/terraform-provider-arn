# arn:aws:amplify:ap-northeast-1:111111111111:apps/app-id/branches/branch-name
output "amplify_branches" {
  value = provider::arn::amplify_branches("app-id", "branch-name")
}
