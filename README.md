REG.RU module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [reg.ru](https://reg.ru).

## Caddy module name

```
dns.providers.regru
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "regru",
				"username": "<USERNAME>",
				"password": "<PASSWORD>"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns regru {
	    username <username>
	    password <password>
	}
}
```

```
# one site
tls {
	dns regru {
	    username <username>
	    password <password>
	}
}
```
