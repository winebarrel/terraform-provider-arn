# arn:aws:apptest:ap-northeast-1:111111111111:testsuite/test-suite-id
output "apptest_test_suite" {
  value = provider::arn::apptest_test_suite("test-suite-id")
}
