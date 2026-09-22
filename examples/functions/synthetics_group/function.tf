# arn:aws:synthetics:ap-northeast-1:111111111111:group:group-id
output "synthetics_group" {
  value = provider::arn::synthetics_group("group-id")
}
