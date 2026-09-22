# arn:aws:organizations::111111111111:transfer/o-organization-id/transfer-type/transfer-direction/rt-responsibility-transfer-id
output "organizations_responsibilitytransfer" {
  value = provider::arn::organizations_responsibilitytransfer("organization-id", "transfer-type", "transfer-direction", "responsibility-transfer-id")
}
