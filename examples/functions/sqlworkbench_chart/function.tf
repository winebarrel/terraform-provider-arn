# arn:aws:sqlworkbench:ap-northeast-1:111111111111:chart/resource-id
output "sqlworkbench_chart" {
  value = provider::arn::sqlworkbench_chart("resource-id")
}
