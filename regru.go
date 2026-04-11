package regru

import (
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsregru "github.com/heinwol/libdns-regru/pkg"
)

// Just a small wrapper
type Provider struct{ *libdnsregru.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.regru",
		New: func() caddy.Module { return &Provider{new(libdnsregru.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.Username = repl.ReplaceAll(p.Provider.Username, "")
	p.Provider.Password = repl.ReplaceAll(p.Provider.Password, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	regru {
//	    username <username>
//	    password <password>
//	}
//
// or inline:
//
//	regru <username> <password>
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// for d.Next() {
	d.Next() // consume directive name

	// inline: regru <username> <password>
	if d.NextArg() {
		p.Provider.Username = d.Val()
		if !d.NextArg() {
			return d.ArgErr()
		}
		p.Provider.Password = d.Val()
	} else {
		// block form, see description
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "username":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.Username = strings.TrimSpace(d.Val())
			case "password":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.Password = strings.TrimSpace(d.Val())
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if d.NextArg() {
		return d.Errf("unexpected argument '%s'", d.Val())
	}
	if p.Provider.Username == "" {
		return d.Err("missing username")
	}
	if p.Provider.Password == "" {
		return d.Err("missing password")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
