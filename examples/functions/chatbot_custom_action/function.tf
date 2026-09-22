# arn:aws:chatbot::111111111111:custom-action/action-name
output "chatbot_custom_action" {
  value = provider::arn::chatbot_custom_action("action-name")
}
