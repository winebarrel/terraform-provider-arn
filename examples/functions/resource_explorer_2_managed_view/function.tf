# arn:aws:resource-explorer-2:ap-northeast-1:111111111111:managed-view/managed-view-name/managed-view-uuid
output "resource_explorer_2_managed_view" {
  value = provider::arn::resource_explorer_2_managed_view("managed-view-name", "managed-view-uuid")
}
