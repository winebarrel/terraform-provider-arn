# arn:aws:controltower:ap-northeast-1:111111111111:enabledbaseline/enabled-baseline-id
output "controltower_enabled_baseline" {
  value = provider::arn::controltower_enabled_baseline("enabled-baseline-id")
}
