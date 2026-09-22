# arn:aws:workspaces:ap-northeast-1:111111111111:workspacecertificate/certificate-id
output "workspaces_certificateid" {
  value = provider::arn::workspaces_certificateid("certificate-id")
}
