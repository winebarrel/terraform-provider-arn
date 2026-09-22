# arn:aws:chatbot::111111111111:chat-configuration/configuration-type/chatbot-configuration-name
output "chatbot_chatbot_configuration" {
  value = provider::arn::chatbot_chatbot_configuration("configuration-type", "chatbot-configuration-name")
}
