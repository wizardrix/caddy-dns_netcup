package netcup

import (
	netcup "github.com/wizardrix/libdns_netcup"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *netcup.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.netcup",
		New: func() caddy.Module { return &Provider{new(netcup.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	replacer := caddy.NewReplacer()
	p.Provider.CustomerNumber = replacer.ReplaceAll(p.Provider.CustomerNumber, "")
	p.Provider.APIKey = replacer.ReplaceAll(p.Provider.APIKey, "")
	p.Provider.APIPassword = replacer.ReplaceAll(p.Provider.APIPassword, "")
	return nil
}

// TODO: This is just an example. Update accordingly.
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// providername [<api_token>] {
//     api_token <api_token>
// }
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// for d.Next() {
	// 	if d.NextArg() {
	// 		p.Provider.APIToken = d.Val()
	// 	}
	// 	if d.NextArg() {
	// 		return d.ArgErr()
	// 	}
	// 	for nesting := d.Nesting(); d.NextBlock(nesting); {
	// 		switch d.Val() {
	// 		case "api_token":
	// 			if p.Provider.APIToken != "" {
	// 				return d.Err("API token already set")
	// 			}
	// 			if d.NextArg() {
	// 				p.Provider.APIToken = d.Val()
	// 			}
	// 			if d.NextArg() {
	// 				return d.ArgErr()
	// 			}
	// 		default:
	// 			return d.Errf("unrecognized subdirective '%s'", d.Val())
	// 		}
	// 	}
	// }
	// if p.Provider.APIToken == "" {
	// 	return d.Err("missing API token")
	// }
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
