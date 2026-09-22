# arn:aws:comprehend:ap-northeast-1:111111111111:events-detection-job/job-id
output "comprehend_events_detection_job" {
  value = provider::arn::comprehend_events_detection_job("job-id")
}
