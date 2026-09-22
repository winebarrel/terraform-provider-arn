# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxDatabase/kx-database
output "finspace_kx_database" {
  value = provider::arn::finspace_kx_database("environment-id", "kx-database")
}
