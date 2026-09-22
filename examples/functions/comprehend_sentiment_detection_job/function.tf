# arn:aws:comprehend:ap-northeast-1:111111111111:sentiment-detection-job/job-id
output "comprehend_sentiment_detection_job" {
  value = provider::arn::comprehend_sentiment_detection_job("job-id")
}
