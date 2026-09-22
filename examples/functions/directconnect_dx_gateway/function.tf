# arn:aws:directconnect::111111111111:dx-gateway/direct-connect-gateway-id
output "directconnect_dx_gateway" {
  value = provider::arn::directconnect_dx_gateway("direct-connect-gateway-id")
}
