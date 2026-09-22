# arn:aws:lex:ap-northeast-1:111111111111:intent:intent-name:intent-version
output "lex_intent_version" {
  value = provider::arn::lex_intent_version("intent-name", "intent-version")
}
