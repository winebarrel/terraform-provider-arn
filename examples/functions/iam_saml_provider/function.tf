# arn:aws:iam::111111111111:saml-provider/saml-provider-name
output "iam_saml_provider" {
  value = provider::arn::iam_saml_provider("saml-provider-name")
}
