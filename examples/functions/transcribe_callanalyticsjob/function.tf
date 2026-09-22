# arn:aws:transcribe:ap-northeast-1:111111111111:analytics/job-name
output "transcribe_callanalyticsjob" {
  value = provider::arn::transcribe_callanalyticsjob("job-name")
}
