# arn:aws:globalaccelerator::111111111111:attachment/resource-id
output "globalaccelerator_attachment" {
  value = provider::arn::globalaccelerator_attachment("resource-id")
}
