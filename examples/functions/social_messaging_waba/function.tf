# arn:aws:social-messaging:ap-northeast-1:111111111111:waba/waba-id
output "social_messaging_waba" {
  value = provider::arn::social_messaging_waba("waba-id")
}
