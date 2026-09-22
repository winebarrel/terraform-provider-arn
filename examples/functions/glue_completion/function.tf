# arn:aws:glue:ap-northeast-1:111111111111:completion/completion-id
output "glue_completion" {
  value = provider::arn::glue_completion("completion-id")
}
