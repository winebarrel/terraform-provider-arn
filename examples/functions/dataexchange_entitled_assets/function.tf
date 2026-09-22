# arn:aws:dataexchange:ap-northeast-1::data-sets/data-set-id/revisions/revision-id/assets/asset-id
output "dataexchange_entitled_assets" {
  value = provider::arn::dataexchange_entitled_assets("data-set-id", "revision-id", "asset-id")
}
