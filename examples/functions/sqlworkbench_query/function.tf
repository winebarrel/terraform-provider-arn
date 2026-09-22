# arn:aws:sqlworkbench:ap-northeast-1:111111111111:query/resource-id
output "sqlworkbench_query" {
  value = provider::arn::sqlworkbench_query("resource-id")
}
