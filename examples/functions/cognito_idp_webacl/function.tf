# arn:aws:wafv2:ap-northeast-1:111111111111:scope/webacl/name/id
output "cognito_idp_webacl" {
  value = provider::arn::cognito_idp_webacl("scope", "name", "id")
}
