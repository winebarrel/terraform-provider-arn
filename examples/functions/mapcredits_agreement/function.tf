# arn:aws:mapcredits:::agreement/agreement-id
output "mapcredits_agreement" {
  value = provider::arn::mapcredits_agreement("agreement", "agreement-id")
}
