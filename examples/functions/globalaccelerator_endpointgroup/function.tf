# arn:aws:globalaccelerator::111111111111:accelerator/resource-id/listener/listener-id/endpoint-group/endpoint-group-id
output "globalaccelerator_endpointgroup" {
  value = provider::arn::globalaccelerator_endpointgroup("resource-id", "listener-id", "endpoint-group-id")
}
