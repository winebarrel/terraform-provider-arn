# arn:aws:waf::111111111111:sqlinjectionset/id
output "waf_sqlinjectionmatchset" {
  value = provider::arn::waf_sqlinjectionmatchset("id")
}
