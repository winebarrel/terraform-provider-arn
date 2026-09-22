# arn:aws:appstream:ap-northeast-1:111111111111:app-block-builder/app-block-builder-name
output "appstream_app_block_builder" {
  value = provider::arn::appstream_app_block_builder("app-block-builder-name")
}
