# arn:aws:globalaccelerator::111111111111:accelerator/resource-id/listener/listener-id
output "globalaccelerator_listener" {
  value = provider::arn::globalaccelerator_listener("resource-id", "listener-id")
}
