# Test case for gocloud.dev and sso sessions

Test case for a problem we're having with gocloud.dev and session-based SSO for AWS
CLI.  The session-based profile `usesession` fails in all cases except gocloud v0.29
using AWS Go SDK v2.

```
===== testing profile standalone gocloud-version 0.27 awssdkver v1
Bucket opened ok
This is a test file

===== testing profile standalone gocloud-version 0.29 awssdkver v1
Bucket opened ok
This is a test file

===== testing profile standalone gocloud-version 0.27 awssdkver v2
Bucket opened ok
This is a test file

===== testing profile standalone gocloud-version 0.29 awssdkver v2
Bucket opened ok
This is a test file

===== testing profile usesession gocloud-version 0.27 awssdkver v1
unable to open bucket s3://duanetestfoo?awssdk=v1: open bucket s3://duanetestfoo?awssdk=v1: couldn't create session profile "usesession" is configured to use SSO but is missing required configuration: sso_region, sso_start_url

===== testing profile usesession gocloud-version 0.29 awssdkver v1
unable to open bucket s3://duanetestfoo?awssdk=v1: open bucket s3://duanetestfoo?awssdk=v1: couldn't create session profile "usesession" is configured to use SSO but is missing required configuration: sso_region, sso_start_url

===== testing profile usesession gocloud-version 0.27 awssdkver v2
unable to open bucket s3://duanetestfoo?awssdk=v2: open bucket s3://duanetestfoo?awssdk=v2: profile "usesession" is configured to use SSO but is missing required configuration: sso_region, sso_start_url

===== testing profile usesession gocloud-version 0.29 awssdkver v2
Bucket opened ok
This is a test file
```

# Related

 - https://github.com/aws/aws-sdk-go/issues/4649
