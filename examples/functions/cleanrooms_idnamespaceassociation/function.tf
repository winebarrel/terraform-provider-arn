# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/idnamespaceassociation/id-namespace-association-id
output "cleanrooms_idnamespaceassociation" {
  value = provider::arn::cleanrooms_idnamespaceassociation("membership-id", "id-namespace-association-id")
}
