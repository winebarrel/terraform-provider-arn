# arn:aws:comprehend:ap-northeast-1:111111111111:flywheel/flywheel-name
output "comprehend_flywheel" {
  value = provider::arn::comprehend_flywheel("flywheel-name")
}
