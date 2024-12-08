│ Error: Invalid provider configuration
│
│ Provider "registry.terraform.io/grafana/grafana" requires explicit configuration. Add a provider block to the root module and configure the provider's required
│ arguments as described in the provider documentation.
│
╵
╷
│ Error: Missing required argument
│
│   with provider["registry.terraform.io/grafana/grafana"],
│   on <empty> line 0:
│   (source code not available)
│
│ "auth": one of `auth,cloud_api_key,oncall_access_token,sm_access_token` must be specified
