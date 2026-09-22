# arn:aws:lex:ap-northeast-1:111111111111:bot:bot-name:bot-alias
output "lex_bot_alias_2" {
  value = provider::arn::lex_bot_alias_2("bot-name", "bot-alias")
}
