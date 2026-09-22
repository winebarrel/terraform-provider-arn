# arn:aws:wisdom:ap-northeast-1:111111111111:message-template/knowledge-base-id/message-template-id
output "wisdom_message_template" {
  value = provider::arn::wisdom_message_template("knowledge-base-id", "message-template-id")
}
