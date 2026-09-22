# arn:aws:sqlworkbench:ap-northeast-1:111111111111:connection/resource-id
output "sqlworkbench_connection" {
  value = provider::arn::sqlworkbench_connection("resource-id")
}
