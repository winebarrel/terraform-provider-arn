# arn:aws:rds:ap-northeast-1:111111111111:cev:engine/engine-version/custom-db-engine-version-id
output "rds_cev" {
  value = provider::arn::rds_cev("engine", "engine-version", "custom-db-engine-version-id")
}
