# arn:aws:comprehend:ap-northeast-1:111111111111:flywheel/flywheel-name/dataset/dataset-name
output "comprehend_flywheel_dataset" {
  value = provider::arn::comprehend_flywheel_dataset("flywheel-name", "dataset-name")
}
