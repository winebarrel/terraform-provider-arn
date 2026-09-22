# arn:aws:fis:ap-northeast-1:111111111111:experiment/id
output "fis_experiment" {
  value = provider::arn::fis_experiment("id")
}
