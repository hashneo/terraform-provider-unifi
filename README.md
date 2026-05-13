# terraform-provider-unifi

A custom Terraform provider for Unifi network infrastructure. Supports both the **Unifi Site Manager cloud API** (`api.ui.com`) and local controller APIs. Covers Network, Protect (cameras/sensors), and Site Manager endpoints.

Tested against: Unifi cloud-managed network, Network application v10.3.58, UniFi Protect.

---

## Requirements

- [Go](https://golang.org/) >= 1.21
- [Terraform](https://www.terraform.io/) >= 1.5.0
- macOS: `codesign` (included with Xcode Command Line Tools)
- A Unifi **cloud API key** (Site Manager → API Keys) or a local controller API key

---

## Building and installing

```bash
make install
```

This builds the provider binary and copies it to `~/.terraform.d/plugins/registry.terraform.io/local/unifi/unifi/0.1.0/<os>_<arch>/`.

On **macOS** you must sign the binary after install to avoid Gatekeeper killing it:

```bash
codesign --sign - ~/.terraform.d/plugins/registry.terraform.io/local/unifi/unifi/0.1.0/darwin_arm64/terraform-provider-unifi
```

---

## Terraform configuration

### `~/.terraformrc` — dev overrides

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/local/unifi/unifi" = "/Users/<you>/.terraform.d/plugins/registry.terraform.io/local/unifi/unifi/0.1.0/darwin_arm64"
  }
  direct {}
}
```

### `required_providers` block

```hcl
terraform {
  required_version = ">= 1.5.0"
  required_providers {
    unifi = {
      source  = "registry.terraform.io/local/unifi/unifi"
      version = "0.1.0"
    }
  }
}

# Cloud API (Site Manager) — recommended
provider "unifi" {
  cloud_api_key = "your-cloud-api-key"
  site_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}

# Local controller (alternative)
# provider "unifi" {
#   controller_url = "https://10.0.1.1"
#   local_api_key  = "your-local-api-key"
#   site_id        = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#   insecure       = true
# }
```

### Finding your site ID

The site ID is a UUID visible in the Unifi console URL or via the Site Manager API. It looks like `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.

---

## Data sources

### Site Manager

| Data source | Description |
|---|---|
| `unifi_hosts` | All hosts (consoles) registered to your Site Manager account. |
| `unifi_sites` | All sites across your consoles. |
| `unifi_cloud_devices` | All devices visible in Site Manager with adoption/connection state. |
| `unifi_sdwan_configs` | SD-WAN configurations. |

### Network

| Data source | Description |
|---|---|
| `unifi_adopted_devices` | All adopted network devices (APs, switches, gateways) with firmware version and features. |
| `unifi_pending_devices` | Devices discovered but not yet adopted. |
| `unifi_clients` | All connected network clients with MAC, IP, type, and uplink device. |
| `unifi_networks` | Configured networks / VLANs. |
| `unifi_wifi_broadcasts` | Wi-Fi SSIDs with security type and network assignment. |
| `unifi_firewall_zones` | Zone-based firewall zones (Internal, External, DMZ, VPN, Hotspot, Gateway). |
| `unifi_firewall_policies` | Zone-based firewall policies with source/destination zones, IP scope, logging. |
| `unifi_acl_rules` | ACL rules. |
| `unifi_dns_policies` | DNS policies. |
| `unifi_vpn_servers` | VPN server configurations. |
| `unifi_wan_interfaces` | WAN interface list. |
| `unifi_device_tags` | Device tag groups with member device IDs. |

### Protect

| Data source | Description |
|---|---|
| `unifi_nvr` | NVR (Network Video Recorder) identity and model info. |
| `unifi_cameras` | Camera inventory with state and microphone status. |
| `unifi_sensors` | Door/motion sensors with battery level, motion state, open/closed state. |
| `unifi_lights` | Unifi floodlights with PIR motion state. |
| `unifi_viewers` | Unifi Viewer devices. |

---

## Resources

| Resource | Description |
|---|---|
| `unifi_network` | Manage a network / VLAN. |
| `unifi_wifi` | Manage a Wi-Fi SSID. |
| `unifi_firewall_policy` | Manage a zone-based firewall policy. |
| `unifi_acl_rule` | Manage an ACL rule. |

---

## unifi-report — standalone HTML inventory report

A standalone tool that connects to your Unifi console and produces a self-contained HTML (or JSON) inventory report without requiring Terraform.

### Run with cloud API key

```bash
go run ./cmd/unifi-report/ \
  --cloud-key <your-cloud-api-key> \
  --site-id   xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  --out       unifi-report.html
```

### Run against a local controller

```bash
go run ./cmd/unifi-report/ \
  --controller https://192.168.1.1 \
  --local-key  your-local-api-key \
  --site-id    xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  --insecure \
  --out        unifi-report.html
```

### Environment variables

```bash
export UNIFI_CLOUD_API_KEY=<your-cloud-api-key>
export UNIFI_SITE_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

go run ./cmd/unifi-report/ --out unifi-report.html
```

### Flags

| Flag | Env var | Default | Description |
|---|---|---|---|
| `--cloud-key` | `UNIFI_CLOUD_API_KEY` | — | Site Manager / cloud API key |
| `--local-key` | `UNIFI_LOCAL_API_KEY` | — | Local controller API key |
| `--controller` | `UNIFI_CONTROLLER_URL` | — | Local controller URL, e.g. `https://192.168.1.1` |
| `--site-id` | `UNIFI_SITE_ID` | — | Site UUID (required for Network and Protect endpoints) |
| `--out` | — | `unifi-report.html` | Output file. Use `.json` extension for JSON output |
| `--insecure` | — | `false` | Skip TLS certificate verification |

### Output

The report covers 21 sections across 3 groups:

| Group | Sections |
|---|---|
| Site Manager | Hosts, Sites, Cloud Devices, SD-WAN Configs |
| Network | Adopted Devices, Pending Devices, Clients, Networks, Wi-Fi SSIDs, Firewall Zones, Firewall Policies, ACL Rules, DNS Policies, VPN Servers, WAN Interfaces, Device Tags |
| Protect | NVR, Cameras, Sensors, Lights, Viewers |

All sections are fetched in parallel. A full report typically completes in under 10 seconds.

### JSON output

```bash
go run ./cmd/unifi-report/ --cloud-key <key> --site-id <id> --out unifi-report.json
```

---

## API architecture

All network and protect data is proxied through the Unifi Site Manager cloud API:

```
https://api.ui.com/v1/connector/consoles/{consoleId}/proxy/network/...
https://api.ui.com/v1/connector/consoles/{consoleId}/proxy/protect/...
```

The client performs a one-time `discover()` call on startup to resolve the console ID from the site ID, then routes all subsequent requests through the proxy.

List endpoints are paginated using a generic `listAllNetwork[T]` helper that fetches in pages of 100 until all items are retrieved.

## Known limitations

- **Zone-based firewall only.** The provider targets the modern zone-based firewall API (UniFi Network 8+). Legacy rule-based firewall is not supported.
- **Cloud proxy required for Protect.** UniFi Protect data is only accessible via the cloud proxy; direct local Protect API access is not implemented.
- **Read-mostly.** Write operations are implemented for Network, WiFi, Firewall Policy, and ACL Rule resources. Most other resources are read-only data sources.
