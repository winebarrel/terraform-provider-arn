# arn:aws:comprehend:ap-northeast-1:111111111111:entities-detection-job/job-id
output "comprehend_entities_detection_job" {
  value = provider::arn::comprehend_entities_detection_job("job-id")
}
