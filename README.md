servedir
========

A trivial Go (golang) application to serve a directory over HTTP or HTTPS.


usage
-----

```
usage: servedir [<flags>]

Serve a directory over HTTP(S)

Flags:
  --help              Show help (also see --help-long and --help-man).
  -a, --addr=":8080"  Address to listen on
  -d, --dir="."       Directory to serve
  --cert=CERT         TLS certificate to use (optional)
  --key=KEY           TLS certificate private key to use when using --cert
  --version           Show application version.
```
