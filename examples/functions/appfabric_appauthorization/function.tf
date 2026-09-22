# arn:aws:appfabric:ap-northeast-1:111111111111:appbundle/appbundle-id/appauthorization/app-authorization-identifier
output "appfabric_appauthorization" {
  value = provider::arn::appfabric_appauthorization("appbundle-id", "app-authorization-identifier")
}
