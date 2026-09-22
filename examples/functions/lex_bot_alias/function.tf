# arn:aws:lex:ap-northeast-1:111111111111:bot-alias/bot-id/bot-alias-id
output "lex_bot_alias" {
  value = provider::arn::lex_bot_alias("bot-id", "bot-alias-id")
}
