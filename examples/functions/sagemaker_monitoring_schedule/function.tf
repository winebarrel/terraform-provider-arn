# arn:aws:sagemaker:ap-northeast-1:111111111111:monitoring-schedule/monitoring-schedule-name
output "sagemaker_monitoring_schedule" {
  value = provider::arn::sagemaker_monitoring_schedule("monitoring-schedule-name")
}
