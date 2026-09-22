# arn:aws:glue:ap-northeast-1:111111111111:trigger/trigger-name
output "glue_trigger" {
  value = provider::arn::glue_trigger("trigger-name")
}
