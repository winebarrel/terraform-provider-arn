# arn:aws:pca-connector-ad:ap-northeast-1:111111111111:connector/connector-id/template/template-id
output "pca_connector_ad_template" {
  value = provider::arn::pca_connector_ad_template("connector-id", "template-id")
}
