# arn:aws:pca-connector-scep:ap-northeast-1:111111111111:connector/connector-id
output "pca_connector_scep_connector" {
  value = provider::arn::pca_connector_scep_connector("connector-id")
}
