# arn:aws:rolesanywhere:ap-northeast-1:111111111111:crl/crl-id
output "rolesanywhere_crl" {
  value = provider::arn::rolesanywhere_crl("crl-id")
}
