# arn:aws:directconnect::111111111111:dx-resiliency-group/resiliency-group-id
output "directconnect_dx_resiliency_group" {
  value = provider::arn::directconnect_dx_resiliency_group("resiliency-group-id")
}
