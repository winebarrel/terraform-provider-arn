# arn:aws:fms:ap-northeast-1:111111111111:applications-list/id
output "fms_applications_list" {
  value = provider::arn::fms_applications_list("id")
}
