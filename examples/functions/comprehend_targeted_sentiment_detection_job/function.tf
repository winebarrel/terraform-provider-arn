# arn:aws:comprehend:ap-northeast-1:111111111111:targeted-sentiment-detection-job/job-id
output "comprehend_targeted_sentiment_detection_job" {
  value = provider::arn::comprehend_targeted_sentiment_detection_job("job-id")
}
