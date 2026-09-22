# arn:aws:lex:ap-northeast-1:111111111111:bot:bot-name:bot-version
output "lex_bot_version" {
  value = provider::arn::lex_bot_version("bot-name", "bot-version")
}
