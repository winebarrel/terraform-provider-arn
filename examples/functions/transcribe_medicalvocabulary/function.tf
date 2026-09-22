# arn:aws:transcribe:ap-northeast-1:111111111111:medical-vocabulary/vocabulary-name
output "transcribe_medicalvocabulary" {
  value = provider::arn::transcribe_medicalvocabulary("vocabulary-name")
}
