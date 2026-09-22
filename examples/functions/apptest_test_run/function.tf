# arn:aws:apptest:ap-northeast-1:111111111111:testrun/test-run-id
output "apptest_test_run" {
  value = provider::arn::apptest_test_run("test-run-id")
}
