# arn:aws:ivs:ap-northeast-1:111111111111:ingest-configuration/resource-id
output "ivs_ingest_configuration" {
  value = provider::arn::ivs_ingest_configuration("resource-id")
}
