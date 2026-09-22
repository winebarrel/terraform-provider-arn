# arn:aws:appfabric:ap-northeast-1:111111111111:appbundle/app-bundle-identifier
output "appfabric_appbundle" {
  value = provider::arn::appfabric_appbundle("app-bundle-identifier")
}
