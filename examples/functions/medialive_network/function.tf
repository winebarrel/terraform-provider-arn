# arn:aws:medialive:ap-northeast-1:111111111111:network:network-id
output "medialive_network" {
  value = provider::arn::medialive_network("network-id")
}
