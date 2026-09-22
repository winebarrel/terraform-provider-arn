# arn:aws:appstream:ap-northeast-1:111111111111:app-block/app-block-name
output "appstream_app_block" {
  value = provider::arn::appstream_app_block("app-block-name")
}
