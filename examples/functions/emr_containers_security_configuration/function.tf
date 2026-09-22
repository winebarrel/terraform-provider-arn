# arn:aws:emr-containers:ap-northeast-1:111111111111:/securityconfigurations/security-configuration-id
output "emr_containers_security_configuration" {
  value = provider::arn::emr_containers_security_configuration("security-configuration-id")
}
