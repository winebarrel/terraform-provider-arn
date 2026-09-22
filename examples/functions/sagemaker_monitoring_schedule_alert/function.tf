# arn:aws:sagemaker:ap-northeast-1:111111111111:monitoring-schedule/monitoring-schedule-name/alert/monitoring-schedule-alert-name
output "sagemaker_monitoring_schedule_alert" {
  value = provider::arn::sagemaker_monitoring_schedule_alert("monitoring-schedule-name", "monitoring-schedule-alert-name")
}
