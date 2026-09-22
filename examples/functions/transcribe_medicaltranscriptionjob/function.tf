# arn:aws:transcribe:ap-northeast-1:111111111111:medical-transcription-job/job-name
output "transcribe_medicaltranscriptionjob" {
  value = provider::arn::transcribe_medicaltranscriptionjob("job-name")
}
