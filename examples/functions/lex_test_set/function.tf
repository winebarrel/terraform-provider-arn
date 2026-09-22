# arn:aws:lex:ap-northeast-1:111111111111:test-set/test-set-id
output "lex_test_set" {
  value = provider::arn::lex_test_set("test-set-id")
}
