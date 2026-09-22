# arn:aws:ssm-sap:ap-northeast-1:111111111111:application-type/application-id
output "ssm_sap_application" {
  value = provider::arn::ssm_sap_application("application-type", "application-id")
}
