# arn:aws:workspaces-web:ap-northeast-1:111111111111:trustStore/trust-store-id
output "workspaces_web_trust_store" {
  value = provider::arn::workspaces_web_trust_store("trust-store-id")
}
