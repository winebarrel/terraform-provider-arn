# arn:aws:lex:ap-northeast-1:111111111111:bot:bot-name
output "lex_bot_2" {
  value = provider::arn::lex_bot_2("bot-name")
}
