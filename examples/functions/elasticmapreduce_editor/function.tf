# arn:aws:elasticmapreduce:ap-northeast-1:111111111111:editor/editor-id
output "elasticmapreduce_editor" {
  value = provider::arn::elasticmapreduce_editor("editor-id")
}
