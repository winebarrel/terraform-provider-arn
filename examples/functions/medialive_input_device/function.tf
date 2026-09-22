# arn:aws:medialive:ap-northeast-1:111111111111:inputDevice:device-id
output "medialive_input_device" {
  value = provider::arn::medialive_input_device("device-id")
}
