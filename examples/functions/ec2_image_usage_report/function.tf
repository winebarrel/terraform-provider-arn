# arn:aws:ec2:ap-northeast-1:111111111111:image-usage-report/image-usage-report-id
output "ec2_image_usage_report" {
  value = provider::arn::ec2_image_usage_report("image-usage-report-id")
}
