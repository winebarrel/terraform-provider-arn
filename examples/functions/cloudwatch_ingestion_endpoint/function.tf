# arn:aws:cloudwatch:ap-northeast-1:111111111111:ingestion-endpoint/ingestion-endpoint-name/ingestion-endpoint-id
output "cloudwatch_ingestion_endpoint" {
  value = provider::arn::cloudwatch_ingestion_endpoint("ingestion-endpoint-name", "ingestion-endpoint-id")
}
