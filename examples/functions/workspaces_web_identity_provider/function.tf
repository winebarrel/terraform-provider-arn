# arn:aws:workspaces-web:ap-northeast-1:111111111111:identityProvider/portal-id/identity-provider-id
output "workspaces_web_identity_provider" {
  value = provider::arn::workspaces_web_identity_provider("portal-id", "identity-provider-id")
}
