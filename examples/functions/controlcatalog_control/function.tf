# arn:aws:controlcatalog:::control/control-id
output "controlcatalog_control" {
  value = provider::arn::controlcatalog_control("control-id")
}
