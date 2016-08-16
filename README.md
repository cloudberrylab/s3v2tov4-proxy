# s3v2tov4-proxy

AWS S3 Signature-V2 to Signature-v4 proxy.

This proxy converts signature-V2 requests to signature-V4 requests providing compatibilty between
Minio (which needs Signature-V4) with clients using legacy Signature-V2 mechanism for authentication.

Usage:
```
s3v2tov4-proxy -l :8000 -f http://localhost:9000 -access <ACCESSKEY> -secret <SECRETKEY>
```
