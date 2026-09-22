# arn:aws:dataexchange:ap-northeast-1::data-sets/data-set-id/revisions/revision-id
output "dataexchange_entitled_revisions" {
  value = provider::arn::dataexchange_entitled_revisions("data-set-id", "revision-id")
}
