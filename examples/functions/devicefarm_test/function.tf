# arn:aws:devicefarm:ap-northeast-1:111111111111:test:resource-id
output "devicefarm_test" {
  value = provider::arn::devicefarm_test("resource-id")
}
