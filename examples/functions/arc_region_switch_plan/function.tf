# arn:aws:arc-region-switch::111111111111:plan/resource-id
output "arc_region_switch_plan" {
  value = provider::arn::arc_region_switch_plan("resource-id")
}
