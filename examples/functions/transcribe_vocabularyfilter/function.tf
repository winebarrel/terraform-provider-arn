# arn:aws:transcribe:ap-northeast-1:111111111111:vocabulary-filter/vocabulary-filter-name
output "transcribe_vocabularyfilter" {
  value = provider::arn::transcribe_vocabularyfilter("vocabulary-filter-name")
}
