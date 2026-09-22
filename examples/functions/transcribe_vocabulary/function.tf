# arn:aws:transcribe:ap-northeast-1:111111111111:vocabulary/vocabulary-name
output "transcribe_vocabulary" {
  value = provider::arn::transcribe_vocabulary("vocabulary-name")
}
