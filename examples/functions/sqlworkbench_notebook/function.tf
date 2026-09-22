# arn:aws:sqlworkbench:ap-northeast-1:111111111111:notebook/resource-id
output "sqlworkbench_notebook" {
  value = provider::arn::sqlworkbench_notebook("resource-id")
}
