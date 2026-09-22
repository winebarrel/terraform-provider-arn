# arn:aws:resource-explorer-2:ap-northeast-1:111111111111:view/view-name/view-uuid
output "resource_explorer_2_view" {
  value = provider::arn::resource_explorer_2_view("view-name", "view-uuid")
}
