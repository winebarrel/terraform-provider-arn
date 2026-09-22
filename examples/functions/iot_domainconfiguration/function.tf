# arn:aws:iot:ap-northeast-1:111111111111:domainconfiguration/domain-configuration-name/id
output "iot_domainconfiguration" {
  value = provider::arn::iot_domainconfiguration("domain-configuration-name", "id")
}
