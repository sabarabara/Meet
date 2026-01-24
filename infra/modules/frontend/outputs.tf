output "kms_key_id" {
  description = "KMS key ID for S3 encryption"
  value       = aws_kms_key.s3.id
}

output "waf_arn" {
  description = "WAF ACL ARN for CloudFront"
  value       = aws_wafv2_web_acl.cloudfront.arn
}
