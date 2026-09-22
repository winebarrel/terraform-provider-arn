# arn:aws:chime:ap-northeast-1:111111111111:sma/sip-media-application-id
output "chime_sip_media_application" {
  value = provider::arn::chime_sip_media_application("sip-media-application-id")
}
