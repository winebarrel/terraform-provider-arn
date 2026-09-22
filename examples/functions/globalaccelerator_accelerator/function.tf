# arn:aws:globalaccelerator::111111111111:accelerator/resource-id
output "globalaccelerator_accelerator" {
  value = provider::arn::globalaccelerator_accelerator("resource-id")
}
