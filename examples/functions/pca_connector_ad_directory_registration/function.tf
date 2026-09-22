# arn:aws:pca-connector-ad:ap-northeast-1:111111111111:directory-registration/directory-id
output "pca_connector_ad_directory_registration" {
  value = provider::arn::pca_connector_ad_directory_registration("directory-id")
}
