# arn:aws:lex:ap-northeast-1:111111111111:bot-channel:bot-name:bot-alias:channel-name
output "lex_channel" {
  value = provider::arn::lex_channel("bot-name", "bot-alias", "channel-name")
}
