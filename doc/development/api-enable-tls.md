## Enable TLS in the Kore API server

These instructions assume you will use 10443 as your HTTPS port. You can still use 10080.

### Prerequisites

Add `https://localhost:10443/oauth/callback` to the allowed callback URL list in your IDP.

### When using Helm

You have to set the following Helm variables in `my_values.yaml`:

```
api:
  endpoint:
    url: https://localhost:10443
  port: 10443
  hostPort: 10443
  tls:
    enabled: true
```

### When using local up

```
kore alpha local up \
  --set="api.endpoint.url=https://localhost:10443" \
  --set="api.port=10443" \
  --set="api.hostPort=10443"\
   --set="api.tls.enabled=true"
```

### Setting the CA certificate in the Kore CLI configuration

Run `kore login` or `kore profiles configure` to automatically import a local CA certificate

E.g.:
```
kore login -a https://localhost:10443 local -f
```

The CA certificate will be stored in `.kore/config` with the server configuration.
