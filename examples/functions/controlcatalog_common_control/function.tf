# arn:aws:controlcatalog:::common-control/common-control-id
output "controlcatalog_common_control" {
  value = provider::arn::controlcatalog_common_control("common-control-id")
}
