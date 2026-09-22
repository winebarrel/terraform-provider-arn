# arn:aws:resiliencehub:ap-northeast-1:111111111111:test-template/test-template-id
output "resiliencehub_test_template" {
  value = provider::arn::resiliencehub_test_template("test-template-id")
}
