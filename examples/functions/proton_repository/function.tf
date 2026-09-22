# arn:aws:proton:ap-northeast-1:111111111111:repository/provider:name
output "proton_repository" {
  value = provider::arn::proton_repository("provider", "name")
}
