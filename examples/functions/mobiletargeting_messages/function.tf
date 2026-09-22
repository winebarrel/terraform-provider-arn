# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/messages
output "mobiletargeting_messages" {
  value = provider::arn::mobiletargeting_messages("app-id")
}
