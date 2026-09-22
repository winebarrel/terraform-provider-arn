# arn:aws:transcribe:ap-northeast-1:111111111111:medical-scribe-job/job-name
output "transcribe_medicalscribejob" {
  value = provider::arn::transcribe_medicalscribejob("job-name")
}
