# arn:aws:networkmanager::111111111111:attachment/resource-id
output "networkmanager_attachment" {
  value = provider::arn::networkmanager_attachment("resource-id")
}
