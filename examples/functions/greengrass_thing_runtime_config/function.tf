# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/things/thing-name/runtimeconfig
output "greengrass_thing_runtime_config" {
  value = provider::arn::greengrass_thing_runtime_config("thing-name")
}
