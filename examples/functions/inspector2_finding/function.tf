# arn:aws:inspector2:ap-northeast-1:111111111111:finding/finding-id
output "inspector2_finding" {
  value = provider::arn::inspector2_finding("finding-id")
}
