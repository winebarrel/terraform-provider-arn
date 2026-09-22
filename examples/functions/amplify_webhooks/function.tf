# arn:aws:amplify:ap-northeast-1:111111111111:webhooks/webhook-id
output "amplify_webhooks" {
  value = provider::arn::amplify_webhooks("webhook-id")
}
