# arn:aws:ec2:ap-northeast-1:111111111111:application-status-check/application-status-check-id
output "ec2_application_status_check" {
  value = provider::arn::ec2_application_status_check("application-status-check-id")
}
