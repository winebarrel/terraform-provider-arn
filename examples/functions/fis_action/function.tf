# arn:aws:fis:ap-northeast-1:111111111111:action/id
output "fis_action" {
  value = provider::arn::fis_action("id")
}
