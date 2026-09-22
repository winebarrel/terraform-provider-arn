# arn:aws:drs:ap-northeast-1:111111111111:replication-configuration-template/replication-configuration-template-id
output "drs_replication_configuration_template_resource" {
  value = provider::arn::drs_replication_configuration_template_resource("replication-configuration-template-id")
}
