# arn:aws:transcribe:ap-northeast-1:111111111111:language-model/model-name
output "transcribe_languagemodel" {
  value = provider::arn::transcribe_languagemodel("model-name")
}
