# arn:aws:ecr-public::111111111111:registry/registry-id
output "ecr_public_registry" {
  value = provider::arn::ecr_public_registry("registry-id")
}
