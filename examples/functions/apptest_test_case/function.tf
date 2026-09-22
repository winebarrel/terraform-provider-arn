# arn:aws:apptest:ap-northeast-1:111111111111:testcase/test-case-id
output "apptest_test_case" {
  value = provider::arn::apptest_test_case("test-case-id")
}
