# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/things/thing-name/connectivityInfo
output "greengrass_connectivity_info" {
  value = provider::arn::greengrass_connectivity_info("thing-name")
}
