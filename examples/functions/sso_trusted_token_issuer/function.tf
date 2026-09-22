# arn:aws:sso::111111111111:trustedTokenIssuer/instance-id/trusted-token-issuer-id
output "sso_trusted_token_issuer" {
  value = provider::arn::sso_trusted_token_issuer("instance-id", "trusted-token-issuer-id")
}
