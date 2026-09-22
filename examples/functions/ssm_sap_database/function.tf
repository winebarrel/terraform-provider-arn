# arn:aws:ssm-sap:ap-northeast-1:111111111111:application-type/application-id/DB/database-id
output "ssm_sap_database" {
  value = provider::arn::ssm_sap_database("application-type", "application-id", "database-id")
}
