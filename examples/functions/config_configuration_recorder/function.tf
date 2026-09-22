# arn:aws:config:ap-northeast-1:111111111111:configuration-recorder/recorder-name/recorder-id
output "config_configuration_recorder" {
  value = provider::arn::config_configuration_recorder("recorder-name", "recorder-id")
}
