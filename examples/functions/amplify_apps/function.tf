# arn:aws:amplify:ap-northeast-1:111111111111:apps/app-id
output "amplify_apps" {
  value = provider::arn::amplify_apps("app-id")
}
