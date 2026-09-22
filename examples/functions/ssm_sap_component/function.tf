# arn:aws:ssm-sap:ap-northeast-1:111111111111:application-type/application-id/COMPONENT/component-id
output "ssm_sap_component" {
  value = provider::arn::ssm_sap_component("application-type", "application-id", "component-id")
}
