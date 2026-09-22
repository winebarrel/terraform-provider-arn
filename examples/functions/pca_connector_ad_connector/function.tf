# arn:aws:pca-connector-ad:ap-northeast-1:111111111111:connector/connector-id
output "pca_connector_ad_connector" {
  value = provider::arn::pca_connector_ad_connector("connector-id")
}
