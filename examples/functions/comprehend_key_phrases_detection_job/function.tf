# arn:aws:comprehend:ap-northeast-1:111111111111:key-phrases-detection-job/job-id
output "comprehend_key_phrases_detection_job" {
  value = provider::arn::comprehend_key_phrases_detection_job("job-id")
}
