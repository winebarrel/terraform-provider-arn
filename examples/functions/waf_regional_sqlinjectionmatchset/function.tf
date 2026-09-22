# arn:aws:waf-regional:ap-northeast-1:111111111111:sqlinjectionset/id
output "waf_regional_sqlinjectionmatchset" {
  value = provider::arn::waf_regional_sqlinjectionmatchset("id")
}
