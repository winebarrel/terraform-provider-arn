# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/faq/faq-id
output "kendra_faq" {
  value = provider::arn::kendra_faq("index-id", "faq-id")
}
