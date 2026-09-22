# arn:aws:resource-groups:ap-northeast-1:111111111111:group/group-name/tag-sync-task/task-id
output "resource_groups_tag_sync_task" {
  value = provider::arn::resource_groups_tag_sync_task("group-name", "task-id")
}
