# arn:aws:appfabric:ap-northeast-1:111111111111:appbundle/appbundle-id/ingestion/ingestion-identifier/ingestiondestination/ingestion-destination-identifier
output "appfabric_ingestiondestination" {
  value = provider::arn::appfabric_ingestiondestination("appbundle-id", "ingestion-identifier", "ingestion-destination-identifier")
}
