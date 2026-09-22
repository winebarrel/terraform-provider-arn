# arn:aws:apptest:ap-northeast-1:111111111111:testconfiguration/test-configuration-id
output "apptest_test_configuration" {
  value = provider::arn::apptest_test_configuration("test-configuration-id")
}
