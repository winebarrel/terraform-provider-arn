# arn:aws:appfabric:ap-northeast-1:111111111111:appbundle/appbundle-id/ingestion/ingestion-identifier
output "appfabric_ingestion" {
  value = provider::arn::appfabric_ingestion("appbundle-id", "ingestion-identifier")
}
