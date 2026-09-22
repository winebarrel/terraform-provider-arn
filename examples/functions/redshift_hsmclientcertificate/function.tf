# arn:aws:redshift:ap-northeast-1:111111111111:hsmclientcertificate:hsm-client-certificate-id
output "redshift_hsmclientcertificate" {
  value = provider::arn::redshift_hsmclientcertificate("hsm-client-certificate-id")
}
