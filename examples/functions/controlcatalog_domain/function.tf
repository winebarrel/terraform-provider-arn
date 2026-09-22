# arn:aws:controlcatalog:::domain/domain-id
output "controlcatalog_domain" {
  value = provider::arn::controlcatalog_domain("domain-id")
}
