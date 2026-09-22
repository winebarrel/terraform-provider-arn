# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/users/user-id
output "mobiletargeting_user" {
  value = provider::arn::mobiletargeting_user("app-id", "user-id")
}
