# arn:aws:redshift:ap-northeast-1:111111111111:hsmconfiguration:hsm-configuration-id
output "redshift_hsmconfiguration" {
  value = provider::arn::redshift_hsmconfiguration("hsm-configuration-id")
}
