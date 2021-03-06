# Adding trusted CA certificates

Will add CA certificate to system directories, including Java and Python certificate directories.

```yaml
packages:
  - openssl
  - python3
  - ca-certificates #ubuntu
  - which #+centos
ca:
  - "fixtures/files/example-k8s-ca.pem"
  - "https://dl.cacerts.digicert.com/CybertrustGlobalRoot.crt.pem"
  - |
    -----BEGIN CERTIFICATE-----
    MIIC0jCCAbqgAwIBAgIII8NXV0sAYCAwDQYJKoZIhvcNAQELBQAwFTETMBEGA1UE
    AxMKaW5ncmVzcy1jYTAeFw0yMDAyMTMyMDA1MjdaFw0yMTAyMTEyMDIwMjdaMBUx
    EzARBgNVBAMTCmluZ3Jlc3MtY2EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
    AoIBAQDB25j9ure4vObk+B77G+wbTjpWnGDjCkiR2Ytf9ZRq8Sx+Ee9flwRLe83g
    AKNHhww+5jBACjhx3lKHXACfxOIuIKWw4Ca8CYJc0ydq2aIZgkPzjD+O3lla5Al3
    ZJg4C2bhg+nNG+AepM/Z3V3JIR+ggXzhljp037XJCpJpoHmdQUiEVNRY6bH7gbzP
    WnANYk7yXaVas34JG0qb5y0nwhR6UWqCW1kzJjxGvrmS2itQZa9VSR7d8mXAs34K
    Xm+SNeRWjyu7ivrBAxBP0mAlarUBOMq9mdHN/iL0dd+6ozQUoTqJLiR91lNYku0O
    IWacO0vLYYN1lfXrgdqUluTLii+/AgMBAAGjJjAkMA4GA1UdDwEB/wQEAwICpDAS
    BgNVHRMBAf8ECDAGAQH/AgEAMA0GCSqGSIb3DQEBCwUAA4IBAQCJOaAdIr1lq2qd
    PtNc2wZs3OY4ZMdbIhmrRbC62CPBD/qNIkrl0+ScwmV/Vav6s/jPIdUgtrlKTLzC
    4hsVN7fS4EjLhxJZ4ZNEX6lyhySnQ4H5twVPKVik2A8GabDAfV80BZvtjdrZ0qgP
    Kt+NZcF/5qo7dE/BbQ1w3UWDyE+MkPCGFEBdToc5OqrsiYbHY1bvw8YzSYbB7VEv
    msVQC/+S4F8dFAJ6oLWLnGJ6oIrEJVulIUmeM71syexuQ3G1DS4pcHFjBW7lyjXZ
    WijAm3hUXGPiIwNHIHQ9rW81dLpaS0QUi8l/r4P6Yt1kVrn2UuOSgZ/uf7dC4FdU
    NG5qQLSi
    -----END CERTIFICATE-----
```
