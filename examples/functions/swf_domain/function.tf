# arn:aws:swf::111111111111:/domain/domain-name
output "swf_domain" {
  value = provider::arn::swf_domain("domain-name")
}
