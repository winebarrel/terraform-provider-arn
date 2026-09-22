# arn:aws:wisdom:ap-northeast-1:111111111111:quick-response/knowledge-base-id/quick-response-id
output "wisdom_quick_response" {
  value = provider::arn::wisdom_quick_response("knowledge-base-id", "quick-response-id")
}
