# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/chat-response-configuration/chat-response-configuration-id
output "qbusiness_chat_response_configuration" {
  value = provider::arn::qbusiness_chat_response_configuration("application-id", "chat-response-configuration-id")
}
