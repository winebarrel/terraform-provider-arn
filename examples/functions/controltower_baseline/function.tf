# arn:aws:controltower:ap-northeast-1::baseline/baseline-id
output "controltower_baseline" {
  value = provider::arn::controltower_baseline("baseline-id")
}
