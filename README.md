# migrate-from-apache-to-envoy

```:sh
$ cd certs
$ openssl req -nodes -new -x509 \
  -keyout example-com.key -out example-com.crt \
  -days 365 \
  -subj '/CN=example.com/O=My Company Name LTD./C=US';
$ ls
```
