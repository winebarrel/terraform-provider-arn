# arn:aws:backup:ap-northeast-1:111111111111:tiering-configuration:tiering-configuration-name-tiering-configuration-id
output "backup_tiering_configuration" {
  value = provider::arn::backup_tiering_configuration("tiering-configuration-name", "tiering-configuration-id")
}
