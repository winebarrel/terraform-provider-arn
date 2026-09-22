# arn:aws:glacier:ap-northeast-1:111111111111:vaults/vault-name
output "glacier_vault" {
  value = provider::arn::glacier_vault("vault-name")
}
