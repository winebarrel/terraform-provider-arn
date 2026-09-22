# arn:aws:lex:ap-northeast-1:111111111111:bot/bot-id
output "lex_bot" {
  value = provider::arn::lex_bot("bot-id")
}
