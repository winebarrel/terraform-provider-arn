# arn:aws:athena:ap-northeast-1:111111111111:workgroup/work-group-name
output "athena_workgroup" {
  value = provider::arn::athena_workgroup("work-group-name")
}
