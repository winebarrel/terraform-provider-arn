# arn:aws:acm-pca:ap-northeast-1:111111111111:certificate-authority/certificate-authority-id
output "acm_pca_certificate_authority" {
  value = provider::arn::acm_pca_certificate_authority("certificate-authority-id")
}
