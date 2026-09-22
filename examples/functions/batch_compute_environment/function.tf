# arn:aws:batch:ap-northeast-1:111111111111:compute-environment/compute-environment-name
output "batch_compute_environment" {
  value = provider::arn::batch_compute_environment("compute-environment-name")
}
