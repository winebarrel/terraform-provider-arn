# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id
output "deadline_farm" {
  value = provider::arn::deadline_farm("farm-id")
}
