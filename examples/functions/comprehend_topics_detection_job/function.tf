# arn:aws:comprehend:ap-northeast-1:111111111111:topics-detection-job/job-id
output "comprehend_topics_detection_job" {
  value = provider::arn::comprehend_topics_detection_job("job-id")
}
