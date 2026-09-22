# arn:aws:comprehend:ap-northeast-1:111111111111:dominant-language-detection-job/job-id
output "comprehend_dominant_language_detection_job" {
  value = provider::arn::comprehend_dominant_language_detection_job("job-id")
}
