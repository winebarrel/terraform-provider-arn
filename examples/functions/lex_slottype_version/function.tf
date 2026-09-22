# arn:aws:lex:ap-northeast-1:111111111111:slottype:slot-name:slot-version
output "lex_slottype_version" {
  value = provider::arn::lex_slottype_version("slot-name", "slot-version")
}
