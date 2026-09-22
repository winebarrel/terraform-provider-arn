# arn:aws:transcribe:ap-northeast-1:111111111111:transcription-job/job-name
output "transcribe_transcriptionjob" {
  value = provider::arn::transcribe_transcriptionjob("job-name")
}
