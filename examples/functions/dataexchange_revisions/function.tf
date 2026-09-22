# arn:aws:dataexchange:ap-northeast-1:111111111111:data-sets/data-set-id/revisions/revision-id
output "dataexchange_revisions" {
  value = provider::arn::dataexchange_revisions("data-set-id", "revision-id")
}
