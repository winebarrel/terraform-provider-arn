# arn:aws:opsworks-cm::111111111111:server/server-name/unique-id
output "opsworks_cm_server" {
  value = provider::arn::opsworks_cm_server("server-name", "unique-id")
}
