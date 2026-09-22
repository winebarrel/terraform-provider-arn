# arn:aws:mgn:ap-northeast-1:111111111111:network-migration-definition/network-migration-definition-id
output "mgn_network_migration_definition_resource" {
  value = provider::arn::mgn_network_migration_definition_resource("network-migration-definition-id")
}
