# arn:aws:organizations::111111111111:handshake/o-organization-id/handshake-type/h-handshake-id
output "organizations_handshake" {
  value = provider::arn::organizations_handshake("organization-id", "handshake-type", "handshake-id")
}
