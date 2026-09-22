# arn:aws:controltower:ap-northeast-1:111111111111:enabledcontrol/enabled-control-id
output "controltower_enabled_control" {
  value = provider::arn::controltower_enabled_control("enabled-control-id")
}
