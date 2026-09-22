# arn:aws:rolesanywhere:ap-northeast-1:111111111111:trust-anchor/trust-anchor-id
output "rolesanywhere_trust_anchor" {
  value = provider::arn::rolesanywhere_trust_anchor("trust-anchor-id")
}
