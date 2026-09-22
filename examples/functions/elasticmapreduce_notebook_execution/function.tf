# arn:aws:elasticmapreduce:ap-northeast-1:111111111111:notebook-execution/notebook-execution-id
output "elasticmapreduce_notebook_execution" {
  value = provider::arn::elasticmapreduce_notebook_execution("notebook-execution-id")
}
