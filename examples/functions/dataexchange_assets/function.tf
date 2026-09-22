# arn:aws:dataexchange:ap-northeast-1:111111111111:data-sets/data-set-id/revisions/revision-id/assets/asset-id
output "dataexchange_assets" {
  value = provider::arn::dataexchange_assets("data-set-id", "revision-id", "asset-id")
}
