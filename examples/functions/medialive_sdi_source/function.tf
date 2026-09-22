# arn:aws:medialive:ap-northeast-1:111111111111:sdiSource:sdi-source-id
output "medialive_sdi_source" {
  value = provider::arn::medialive_sdi_source("sdi-source-id")
}
