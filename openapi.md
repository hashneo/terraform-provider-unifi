# UniFi Network API — 10.3.58

## Servers

- `/integration` — 

## Endpoints

### `GET /v1/countries`

**List Countries**

Returns ISO-standard country codes and names,
used for region-based configuration or regulatory compliance.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`code`|`STRING`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/dpi/applications`

**List DPI Applications**

Lists DPI-recognized applications grouped under categories. Useful for firewall or traffic analytics integration.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`INTEGER`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/dpi/categories`

**List DPI Application Categories**

Returns predefined Deep Packet Inspection (DPI) application categories used for traffic identification and filtering.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`INTEGER`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/info`

**Get Application Info**

Retrieve general information about the UniFi Network application.

Tags: Application Info

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/pending-devices`

**List Devices Pending Adoption**

Retrieve a paginated list of devices pending adoption, including basic device information.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`macAddress`|`STRING`|`eq` `ne` `in` `notIn`|
|`ipAddress`|`STRING`|`eq` `ne` `in` `notIn`|
|`model`|`STRING`|`eq` `ne` `in` `notIn`|
|`state`|`STRING`|`eq` `ne` `in` `notIn`|
|`supported`|`BOOLEAN`|`eq` `ne`|
|`firmwareVersion`|`STRING`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le` `like` `in` `notIn`|
|`firmwareUpdatable`|`BOOLEAN`|`eq` `ne`|
|`features`|`SET(STRING)`|`isEmpty` `contains` `containsAny` `containsAll` `containsExactly`|
</details>

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites`

**List Local Sites**

Retrieve a paginated list of local sites managed by this Network application.
Site ID is required for other UniFi Network API calls.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`internalReference`|`STRING`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Sites

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/acl-rules`

**List ACL Rules**

Retrieve a paginated list of all ACL rules on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`enabled`|`BOOLEAN`|`eq` `ne`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`description`|`STRING`|`isNull` `isNotNull` `eq` `ne` `in` `notIn` `like`|
|`action`|`STRING`|`eq` `ne` `in` `notIn`|
|`index`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`protocolFilter`|`SET(STRING)`|`isNull` `isNotNull` `contains` `containsAny` `containsAll` `containsExactly`|
|`networkId`|`UUID`|`isNull` `isNotNull` `eq` `ne` `in` `notIn`|
|`enforcingDeviceFilter.deviceIds`|`SET(UUID)`|`isNull` `isNotNull` `contains` `containsAny` `containsAll` `containsExactly`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
|`sourceFilter.type`|`STRING`|`isNull` `isNotNull` `eq` `ne` `in` `notIn`|
|`sourceFilter.ipAddressesOrSubnets`|`SET(STRING)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`sourceFilter.portFilter`|`SET(INTEGER)`|`isNull` `isNotNull` `contains` `containsAny` `containsAll` `containsExactly`|
|`sourceFilter.networkIds`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`sourceFilter.macAddresses`|`SET(STRING)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`sourceFilter.prefixLength`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`destinationFilter.type`|`STRING`|`isNull` `isNotNull` `eq` `ne` `in` `notIn`|
|`destinationFilter.ipAddressesOrSubnets`|`SET(STRING)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`destinationFilter.portFilter`|`SET(INTEGER)`|`isNull` `isNotNull` `contains` `containsAny` `containsAll` `containsExactly`|
|`destinationFilter.networkIds`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`destinationFilter.macAddresses`|`SET(STRING)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`destinationFilter.prefixLength`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
</details>

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/acl-rules`

**Create ACL Rule**

Create a new user defined ACL rule on a site.

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `ACL rule update`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/acl-rules/ordering`

**Get User-Defined ACL Rule Ordering**

Retrieve user-defined ACL rule ordering on a site.

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/acl-rules/ordering`

**Reorder User-Defined ACL Rules**

Reorder user-defined ACL rules on a site.

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `ACL rule ordering`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/acl-rules/{aclRuleId}`

**Get ACL Rule**

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `aclRuleId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/acl-rules/{aclRuleId}`

**Update ACL Rule**

Update an existing user defined ACL rule on a site.

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `aclRuleId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `ACL rule update`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/acl-rules/{aclRuleId}`

**Delete ACL Rule**

Delete an existing user defined ACL rule on a site.

Tags: Access Control (ACL Rules)

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `aclRuleId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/clients`

**List Connected Clients**

Retrieve a paginated list of all connected clients on a site, including physical devices (computers, smartphones) and active VPN connections.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`macAddress`|`STRING`|`isNull` `isNotNull` `eq` `ne` `in` `notIn`|
|`ipAddress`|`STRING`|`isNull` `isNotNull` `eq` `ne` `in` `notIn`|
|`connectedAt`|`TIMESTAMP`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`access.type`|`STRING`|`eq` `ne` `in` `notIn`|
|`access.authorized`|`BOOLEAN`|`isNull` `isNotNull` `eq` `ne`|
</details>

Tags: Clients

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/clients/{clientId}`

**Get Connected Client Details**

Retrieve detailed information about a specific connected client, including name, IP address, MAC address, connection type and access information.

Tags: Clients

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `clientId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/clients/{clientId}/actions`

**Execute Client Action**

Perform an action on a specific connected client. The request body must include the action name and any applicable input arguments.

Tags: Clients

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `clientId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Client action request`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/device-tags`

**List Device Tags**

Returns all device tags defined within a site, which can be used for WiFi Broadcast assignments.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`deviceIds`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | FilterExpression | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/devices`

**List Adopted Devices**

Retrieve a paginated list of all adopted devices on a site, including basic device information.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`macAddress`|`STRING`|`eq` `ne` `in` `notIn`|
|`ipAddress`|`STRING`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`model`|`STRING`|`eq` `ne` `in` `notIn`|
|`state`|`STRING`|`eq` `ne` `in` `notIn`|
|`supported`|`BOOLEAN`|`eq` `ne`|
|`firmwareVersion`|`STRING`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le` `like` `in` `notIn`|
|`firmwareUpdatable`|`BOOLEAN`|`eq` `ne`|
|`features`|`SET(STRING)`|`isEmpty` `contains` `containsAny` `containsAll` `containsExactly`|
|`interfaces`|`SET(STRING)`|`isEmpty` `contains` `containsAny` `containsAll` `containsExactly`|
</details>

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/devices`

**Adopt Devices**

Adopt a device to a site.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `IntegrationDeviceAdoptionRequestDto`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/devices/{deviceId}`

**Get Adopted Device Details**

Retrieve detailed information about a specific adopted device, including firmware versioning, uplink state, details about device features and interfaces (ports, radios) and other key attributes.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |
| `deviceId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/devices/{deviceId}`

**Remove (Unadopt) Device**

Removes (unadopts) an adopted device from the site. If the device is online, it will be reset to factory defaults.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |
| `deviceId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/devices/{deviceId}/actions`

**Execute Adopted Device Action**

Perform an action on an specific adopted device. The request body must include the action name and any applicable input arguments.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |
| `deviceId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Device action request`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions`

**Execute Port Action**

Perform an action on a specific device port. The request body must include the action name and any applicable input arguments.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `portIdx` | path | integer | Yes |  |
| `siteId` | path | string | Yes |  |
| `deviceId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Port action request`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/devices/{deviceId}/statistics/latest`

**Get Latest Adopted Device Statistics**

Retrieve the latest real-time statistics of a specific adopted device, such as uptime, data transmission rates, CPU and memory utilization.

Tags: UniFi Devices

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |
| `deviceId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/dns/policies`

**List DNS Policies**

Retrieve a paginated list of all DNS policies on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`enabled`|`BOOLEAN`|`eq` `ne`|
|`domain`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`ipv4Address`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`ipv6Address`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`targetDomain`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`mailServerDomain`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`text`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`serverDomain`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`ipAddress`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`ttlSeconds`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`priority`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`service`|`STRING`|`eq` `ne` `in` `notIn`|
|`protocol`|`STRING`|`eq` `ne` `in` `notIn`|
|`port`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`weight`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
</details>

Tags: DNS Policies

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/dns/policies`

**Create DNS Policy**

Create a new DNS policy on a site.

Tags: DNS Policies

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update DNS policy`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/dns/policies/{dnsPolicyId}`

**Get DNS Policy**

Retrieve specific DNS policy.

Tags: DNS Policies

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `dnsPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/dns/policies/{dnsPolicyId}`

**Update DNS Policy**

Update an existing DNS policy on a site.

Tags: DNS Policies

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `dnsPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update DNS policy`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/dns/policies/{dnsPolicyId}`

**Delete DNS Policy**

Delete an existing DNS policy on a site.

Tags: DNS Policies

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `dnsPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/firewall/policies`

**List Firewall Policies**

Retrieve a list of all firewall policies on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`source.zoneId`|`UUID`|`eq` `ne` `in` `notIn`|
|`destination.zoneId`|`UUID`|`eq` `ne` `in` `notIn`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/firewall/policies`

**Create Firewall Policy**

Create a new firewall policy on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update firewall policy`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/firewall/policies/ordering`

**Get User-Defined Firewall Policy Ordering**

Retrieve user-defined firewall policy ordering for a specific source/destination zone pair.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `sourceFirewallZoneId` | query | string | Yes |  |
| `destinationFirewallZoneId` | query | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/firewall/policies/ordering`

**Reorder User-Defined Firewall Policies**

Reorder user-defined firewall policies for a specific source/destination zone pair.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `sourceFirewallZoneId` | query | string | Yes |  |
| `destinationFirewallZoneId` | query | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `IntegrationFirewallPolicyOrderingDto`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/firewall/policies/{firewallPolicyId}`

**Get Firewall Policy**

Retrieve specific firewall policy.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/firewall/policies/{firewallPolicyId}`

**Update Firewall Policy**

Update an existing firewall policy on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update firewall policy`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PATCH /v1/sites/{siteId}/firewall/policies/{firewallPolicyId}`

**Patch Firewall Policy**

Patch an existing firewall policy on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Patch firewall policy`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/firewall/policies/{firewallPolicyId}`

**Delete Firewall Policy**

Delete an existing firewall policy on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallPolicyId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/firewall/zones`

**List Firewall Zones**

Retrieve a list of all firewall zones on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
|`metadata.configurable`|`BOOLEAN`|`eq` `ne` `isNull` `isNotNull`|
</details>

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/firewall/zones`

**Create Custom Firewall Zone**

Create a new custom firewall zone on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update firewall zone`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/firewall/zones/{firewallZoneId}`

**Get Firewall Zone**

Get a firewall zone on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallZoneId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/firewall/zones/{firewallZoneId}`

**Update Firewall Zone**

Update a firewall zone on a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallZoneId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update firewall zone`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/firewall/zones/{firewallZoneId}`

**Delete Custom Firewall Zone**

Delete a custom firewall zone from a site.

Tags: Firewall

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `firewallZoneId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/hotspot/vouchers`

**List Vouchers**

Retrieve a paginated list of Hotspot vouchers.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`createdAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`code`|`STRING`|`eq` `ne` `in` `notIn`|
|`authorizedGuestLimit`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`authorizedGuestCount`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le`|
|`activatedAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`expiresAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`expired`|`BOOLEAN`|`eq` `ne`|
|`timeLimitMinutes`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le`|
|`dataUsageLimitMBytes`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`rxRateLimitKbps`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`txRateLimitKbps`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
</details>

Tags: Hotspot

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/hotspot/vouchers`

**Generate Vouchers**

Create one or more Hotspot vouchers.

Tags: Hotspot

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Hotspot voucher creation request`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `DELETE /v1/sites/{siteId}/hotspot/vouchers`

**Delete Vouchers**

Remove Hotspot vouchers based on the specified filter criteria.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`createdAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`code`|`STRING`|`eq` `ne` `in` `notIn`|
|`authorizedGuestLimit`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`authorizedGuestCount`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le`|
|`activatedAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`expiresAt`|`TIMESTAMP`|`eq` `ne` `gt` `ge` `lt` `le`|
|`expired`|`BOOLEAN`|`eq` `ne`|
|`timeLimitMinutes`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le`|
|`dataUsageLimitMBytes`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`rxRateLimitKbps`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
|`txRateLimitKbps`|`INTEGER`|`isNull` `isNotNull` `eq` `ne` `gt` `ge` `lt` `le`|
</details>

Tags: Hotspot

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `filter` | query | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/hotspot/vouchers/{voucherId}`

**Get Voucher Details**

Retrieve details of a specific Hotspot voucher.

Tags: Hotspot

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `voucherId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/hotspot/vouchers/{voucherId}`

**Delete Voucher**

Remove a specific Hotspot voucher.

Tags: Hotspot

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `voucherId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/networks`

**List Networks**

Retrieve a paginated list of all Networks on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`management`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`enabled`|`BOOLEAN`|`eq` `ne`|
|`vlanId`|`INTEGER`|`eq` `ne` `gt` `ge` `lt` `le` `in` `notIn`|
|`deviceId`|`UUID`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/networks`

**Create Network**

Create a new network on a site.

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update Network`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/networks/{networkId}`

**Get Network Details**

Retrieve detailed information about a specific network.

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `networkId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/networks/{networkId}`

**Update Network**

Update an existing network on a site.

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `networkId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update Network`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/networks/{networkId}`

**Delete Network**

Delete an existing network on a site.

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `networkId` | path | string | Yes |  |
| `force` | query | boolean | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/networks/{networkId}/references`

**Get Network References**

Retrieve references to a specific network.

Tags: Networks

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `networkId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/radius/profiles`

**List Radius Profiles**

Returns available RADIUS authentication profiles, including configuration origin metadata.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/lags`

**List LAGs**

Retrieve a paginated list of all LAGs (Link Aggregation Groups) on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`switchStackId`|`UUID`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
|`mcLagDomainId`|`UUID`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
|`members.deviceId`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`members.portIdxs`|`SET(INTEGER)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/lags/{lagId}`

**Get LAG Details**

Retrieve LAG details.

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `lagId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/mc-lag-domains`

**List MC-LAG Domains**

Retrieve a paginated list of all MC-LAG Domains on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`peers.deviceId`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/mc-lag-domains/{mcLagDomainId}`

**Get MC-LAG Domain**

Retrieve MC-LAG Domain details.

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `mcLagDomainId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/switch-stacks`

**List Switch Stacks**

Retrieve a paginated list of all Switch Stacks on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`members.deviceId`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/switching/switch-stacks/{switchStackId}`

**Get Switch Stack**

Retrieve Switch Stack details.

Tags: Switching

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `switchStackId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/traffic-matching-lists`

**List Traffic Matching Lists**

Retrieve all traffic matching lists on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
</details>

Tags: Traffic Matching Lists

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/traffic-matching-lists`

**Create Traffic Matching List**

Create a new traffic matching list on a site.

Tags: Traffic Matching Lists

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update traffic matching list`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}`

**Get Traffic Matching List**

Get an exist traffic matching list on a site.

Tags: Traffic Matching Lists

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `trafficMatchingListId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}`

**Update Traffic Matching List**

Update an exist traffic matching list on a site.

Tags: Traffic Matching Lists

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `trafficMatchingListId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Create or update traffic matching list`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}`

**Delete Traffic Matching List**

Delete an exist traffic matching list on a site.

Tags: Traffic Matching Lists

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `trafficMatchingListId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/vpn/servers`

**List VPN Servers**

Retrieve a paginated list of all VPN servers on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`enabled`|`BOOLEAN`|`eq` `ne`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/vpn/site-to-site-tunnels`

**List Site-To-Site VPN Tunnels**

Retrieve a paginated list of all site-to-site VPN tunnels on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
|`metadata.source`|`STRING`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
</details>

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/wans`

**List WAN Interfaces**

Returns available WAN interface definitions for a given site,
including identifiers and names. Useful for network and NAT configuration.

Tags: Supporting Resources

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `GET /v1/sites/{siteId}/wifi/broadcasts`

**List Wifi Broadcasts**

Retrieve a paginated list of all Wifi Broadcasts on a site.

<details>
<summary>Filterable properties (click to expand)</summary>

|Name|Type|Allowed functions|
|-|-|-|
|`type`|`STRING`|`eq` `ne` `in` `notIn`|
|`id`|`UUID`|`eq` `ne` `in` `notIn`|
|`enabled`|`BOOLEAN`|`eq` `ne`|
|`name`|`STRING`|`eq` `ne` `in` `notIn` `like`|
|`broadcastingFrequenciesGHz`|`SET(DECIMAL)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`metadata.origin`|`STRING`|`eq` `ne` `in` `notIn`|
|`network.type`|`STRING`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
|`network.networkId`|`UUID`|`eq` `ne` `in` `notIn`|
|`securityConfiguration.type`|`STRING`|`eq` `ne` `in` `notIn`|
|`broadcastingDeviceFilter.type`|`STRING`|`eq` `ne` `in` `notIn` `isNull` `isNotNull`|
|`broadcastingDeviceFilter.deviceIds`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`broadcastingDeviceFilter.deviceTagIds`|`SET(UUID)`|`contains` `containsAny` `containsAll` `containsExactly`|
|`hotspotConfiguration.type`|`STRING`|`eq` `ne` `in` `notIn`|
</details>

Tags: WiFi Broadcasts

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `offset` | query | integer | No |  |
| `limit` | query | integer | No |  |
| `filter` | query | string | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `POST /v1/sites/{siteId}/wifi/broadcasts`

**Create Wifi Broadcast**

Create a new Wifi Broadcast on the specified site.

Tags: WiFi Broadcasts

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Wifi broadcast create or update`

**Responses**

| Status | Description |
|--------|-------------|
| 201 | Created |

---

### `GET /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}`

**Get Wifi Broadcast Details**

Retrieve detailed information about a specific Wifi.

Tags: WiFi Broadcasts

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `wifiBroadcastId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `PUT /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}`

**Update Wifi Broadcast**

Update an existing Wifi Broadcast on the specified site.

Tags: WiFi Broadcasts

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `wifiBroadcastId` | path | string | Yes |  |
| `siteId` | path | string | Yes |  |

**Request Body**

- Content-Type: `application/json` → `Wifi broadcast create or update`

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

### `DELETE /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}`

**Delete Wifi Broadcast**

Delete an existing Wifi Broadcast from the specified site.

Tags: WiFi Broadcasts

**Parameters**

| Name | In | Type | Required | Description |
|------|----|------|----------|-------------|
| `wifiBroadcastId` | path | string | Yes |  |
| `force` | query | boolean | No |  |
| `siteId` | path | string | Yes |  |

**Responses**

| Status | Description |
|--------|-------------|
| 200 | OK |

---

## Schemas

### ACL rule

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `name` | string | Yes | ACL rule name |
| `description` | string |  | ACL rule description |
| `action` | string | Yes | ACL rule action |
| `enforcingDeviceFilter` | ACL rule device filter |  | IDs of the Switch-capable devices used to enforce the ACL rule. When null, the rule will be provisioned to all switches on the site. |
| `index` | integer | Yes | ACL rule index. Lower index has higher priority |
| `sourceFilter` |  |  | Traffic source filter |
| `destinationFilter` |  |  | Traffic destination filter |
| `metadata` | User defined or derived entity metadata | Yes | Only user-defined rules can be deleted or modified |


### ACL rule device filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### ACL rule ordering

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `orderedAclRuleIds` | array | Yes |  |


### ACL rule update

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `name` | string | Yes | ACL rule name |
| `description` | string |  | ACL rule description |
| `action` | string | Yes | ACL rule action |
| `enforcingDeviceFilter` | ACL rule device filter |  | IDs of the Switch-capable devices used to enforce the ACL rule. When null, the rule will be provisioned to all switches on the site. |
| `index` | integer |  | ACL rule index. This property is deprecated and has no effect. Use the dedicated ACL rule reordering endpoint. |
| `sourceFilter` |  |  | Traffic source filter |
| `destinationFilter` |  |  | Traffic destination filter |


### ACL ruleObject

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `name` | string | Yes | ACL rule name |
| `description` | string |  | ACL rule description |
| `action` | string | Yes | ACL rule action |
| `enforcingDeviceFilter` | ACL rule device filter |  | IDs of the Switch-capable devices used to enforce the ACL rule. When null, the rule will be provisioned to all switches on the site. |
| `index` | integer | Yes | ACL rule index. Lower index has higher priority |
| `sourceFilter` |  |  | Traffic source filter |
| `destinationFilter` |  |  | Traffic destination filter |
| `metadata` | User defined or derived entity metadata | Yes | Only user-defined rules can be deleted or modified |


### Access point feature overview


### Address IPv4 matching


### Address IPv6 matching


### Address range IPv4 matching


### Adopted device details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `macAddress` | string | Yes |  |
| `ipAddress` | string | Yes |  |
| `name` | string | Yes |  |
| `model` | string | Yes |  |
| `supported` | boolean | Yes |  |
| `state` | string | Yes |  |
| `firmwareVersion` | string |  |  |
| `firmwareUpdatable` | boolean | Yes |  |
| `adoptedAt` | string |  |  |
| `provisionedAt` | string |  |  |
| `configurationId` | string | Yes |  |
| `uplink` | Device uplink interface overview |  |  |
| `features` | Device features | Yes |  |
| `interfaces` | Device physical interfaces | Yes |  |


### Adopted device overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `macAddress` | string | Yes |  |
| `ipAddress` | string | Yes |  |
| `name` | string | Yes |  |
| `model` | string | Yes |  |
| `state` | string | Yes |  |
| `supported` | boolean | Yes |  |
| `firmwareVersion` | string |  |  |
| `firmwareUpdatable` | boolean | Yes |  |
| `features` | array | Yes |  |
| `interfaces` | array | Yes |  |


### Adopted device overview page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Application info

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `applicationVersion` | string | Yes |  |


### Blackout schedule configuration per day

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `day` | string | Yes |  |


### BooleanType


### Broadcasting device filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Client access overview


### Client action request

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |


### Client action response

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |


### Client details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `connectedAt` | string |  |  |
| `ipAddress` | string |  |  |
| `access` |  | Yes |  |


### Client overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `connectedAt` | string |  |  |
| `ipAddress` | string |  |  |
| `access` | Client access overview | Yes |  |


### Client overview page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### CompoundFilterExpression


### Country Definition

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `code` | string | Yes | The country code in ISO 3166-1 alpha-2 format. |
| `name` | string | Yes | The country name. |


### Country definition page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Create or update DNS policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `enabled` | boolean | Yes |  |


### Create or update Network

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `management` | string | Yes |  |
| `name` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `vlanId` | integer | Yes | VLAN ID. Must be 1 for the default network and >= 2 for additional networks. |
| `dhcpGuarding` | Network DHCP Guarding |  | DHCP Guarding settings for this Network. If this field is omitted or null, the feature is disabled |


### Create or update firewall policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `enabled` | boolean | Yes |  |
| `name` | string | Yes |  |
| `description` | string |  |  |
| `action` | Firewall policy action | Yes |  |
| `source` | Firewall policy source | Yes |  |
| `destination` | Firewall policy destination | Yes |  |
| `ipProtocolScope` | Firewall policy IP protocol scope | Yes |  |
| `connectionStateFilter` | array |  | Match on firewall connection state. If null, matches all connection states. |
| `ipsecFilter` | string |  | Match on traffic encrypted, or not encrypted by IPsec. If null, matches all traffic. |
| `loggingEnabled` | boolean | Yes | Generate syslog entries when traffic is matched. Such entries are sent to a remote syslog server. |
| `schedule` | Firewall schedule |  |  |


### Create or update firewall zone

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | Yes | Name of a firewall zone |
| `networkIds` | array | Yes | List of Network IDs |


### Create or update traffic matching list

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `name` | string | Yes |  |


### DHCP Configuration for IPv6 Network

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ipAddressSuffixRange` | IntegrationIpv6AddressSuffixRangeSelectorDto | Yes |  |
| `leaseTimeSeconds` | integer | Yes | The lease time in seconds for IP addresses in this range. |


### DNS assistance configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | Yes |  |


### DNS policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |
| `domain` | string |  |  |


### DPI application

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | integer | Yes |  |
| `name` | string | Yes |  |


### DPI application page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### DPI category

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | integer | Yes |  |
| `name` | string | Yes |  |


### DPI category page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### DecimalType


### Default client access details


### Default client access overview


### Derived entity metadata


### Device action request

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |


### Device features

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `switching` | Switching feature overview |  |  |
| `accessPoint` | Access point feature overview |  |  |


### Device pending adoption

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `macAddress` | string | Yes |  |
| `ipAddress` | string | Yes |  |
| `model` | string | Yes |  |
| `state` | string | Yes |  |
| `supported` | boolean | Yes |  |
| `firmwareVersion` | string |  |  |
| `firmwareUpdatable` | boolean | Yes |  |
| `features` | array | Yes |  |
| `adoptionTargetSiteIds` | array | Yes |  |


### Device pending adoption page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Device physical interfaces

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ports` | array |  |  |
| `radios` | array |  |  |


### Device restart request


### Device tag

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `deviceIds` | array | Yes |  |
| `metadata` | User or orchestrated entity metadata | Yes |  |


### Device uplink interface overview

Uplink interface is device's connection to the parent device in the network topology

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `deviceId` | string | Yes |  |


### Entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### Error Message

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `statusCode` | integer |  |  |
| `statusName` | string |  |  |
| `code` | string |  |  |
| `message` | string |  |  |
| `timestamp` | string |  |  |
| `requestPath` | string |  |  |
| `requestId` | string |  | In case of Internal Server Error (core = 500), request ID can be used to track down the error in the server log |


### FilterExpression


### FilterPath

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string |  |  |
| `parent` |  |  |  |
| `depth` | integer |  |  |
| `names` | array |  |  |


### FilterableEntity

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `path` | FilterPath |  |  |
| `properties` | object |  |  |
| `nestedEntities` | object |  |  |
| `name` | string |  |  |


### FilterableProperty

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `path` | FilterPath |  |  |
| `type` | FilterablePropertyType |  |  |
| `name` | string |  |  |


### FilterablePropertyType

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `allowedFunctions` | array |  |  |
| `valueType` | string |  |  |
| `supportedFunctions` | array |  |  |


### Firewall policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `name` | string | Yes |  |
| `description` | string |  |  |
| `index` | integer | Yes |  |
| `action` | Firewall policy action | Yes |  |
| `source` | Firewall policy source | Yes |  |
| `destination` | Firewall policy destination | Yes |  |
| `ipProtocolScope` | Firewall policy IP protocol scope | Yes |  |
| `connectionStateFilter` | array |  | Match on firewall connection state. If null, matches all connection states. |
| `ipsecFilter` | string |  | Match on traffic encrypted, or not encrypted by IPsec. If null, matches all traffic. |
| `loggingEnabled` | boolean | Yes | Generate syslog entries when traffic is matched. Such entries are sent to a remote syslog server. |
| `schedule` | Firewall schedule |  |  |
| `metadata` | User or system defined or derived entity metadata | Yes |  |


### Firewall policy IP address filter

Match traffic originating from, or destined to selected IP addresses.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `matchOpposite` | boolean | Yes | Match on all IP addresses except the specified ones. |


### Firewall policy IP protocol scope

Defines rules for matching by IP version and protocol.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ipVersion` | string | Yes |  |


### Firewall policy IPv4 and IPv6 named protocol

Defines rules for matching by protocol name.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | Yes |  |


### Firewall policy IPv4 and IPv6 protocol

Defines protocol matching. If null, matches all protocols.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy IPv4 and IPv6 protocol number

Defines rules for matching by protocol number.


### Firewall policy IPv4 and IPv6 protocol preset

Defines rules for matching by protocol preset.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | Yes |  |


### Firewall policy IPv4 named protocol

Defines rules for matching by protocol name.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string |  |  |


### Firewall policy IPv4 protocol

Defines protocol matching. If null, matches all protocols.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy IPv4 protocol number

Defines rules for matching by protocol number.


### Firewall policy IPv4 protocol preset

Defines rules for matching by protocol preset.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | Yes |  |


### Firewall policy IPv6 interface identifier filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ipv6Iid` | string | Yes | IPv6 Interface Identifier. |
| `matchOpposite` | boolean | Yes | Match on all IPv6 IIDs except the specified one. |


### Firewall policy IPv6 named protocol

Defines rules for matching by protocol name.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string |  |  |


### Firewall policy IPv6 protocol

Defines protocol matching. If null, matches all protocols.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy IPv6 protocol preset

Defines rules for matching by protocol preset.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | Yes |  |


### Firewall policy MAC address filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `macAddresses` | array | Yes | Array of MAC addresses to match. |


### Firewall policy VPN server filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `vpnServerIds` | array | Yes |  |
| `matchOpposite` | boolean | Yes |  |


### Firewall policy action

Defines action for matched traffic.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy application category filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `applicationCategoryIds` | array | Yes | Array of DPI Category IDs to match. |


### Firewall policy application filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `applicationIds` | array | Yes | Array of DPI Application IDs to match. |


### Firewall policy destination

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `zoneId` | string | Yes | ID of the firewall zone to which the matched traffic is destined. |
| `trafficFilter` | Firewall policy destination traffic filter |  |  |


### Firewall policy destination traffic filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy domain filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall policy network filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `networkIds` | array | Yes | Array of Network IDs to match. |
| `matchOpposite` | boolean | Yes | Match on all Networks except the selected. |


### Firewall policy page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Firewall policy port filter

Defines rules for matching traffic by port.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `matchOpposite` | boolean | Yes | Match on all ports except the specified ones. |


### Firewall policy region filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `regions` | array | Yes | Match traffic originating from selected regions. Regions are identified by their ISO 3166-1 alpha-2 country codes. |


### Firewall policy site-to-site VPN tunnel filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `siteToSiteVpnTunnelId` | string | Yes |  |


### Firewall policy source

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `zoneId` | string | Yes | ID of the firewall zone from which the matched traffic originates. |
| `trafficFilter` | Firewall policy source traffic filter |  |  |


### Firewall policy source traffic filter

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Firewall schedule

Defines date and time when the entity is active. If null, the entity is always active.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | Yes |  |


### Firewall schedule time

Defines the time range when the entity is active. If null, the entity is active all day.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `startTime` | string | Yes | Time in HH:MM format. Uses 24-hour clock system. ISO 8601 compliant. |
| `stopTime` | string | Yes | Time in HH:MM format. Uses 24-hour clock system. ISO 8601 compliant. |


### Firewall zone

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes | Name of a firewall zone |
| `networkIds` | array | Yes | List of Network IDs |
| `metadata` | User or system defined entity metadata | Yes | System-defined configurable zones support configuring only attached networks |


### Firewall zones page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Gateway Managed IPv4 Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `autoScaleEnabled` | boolean | Yes | Whether the Network can automatically scale its subnet size based on the number of active DHCP leases. |
| `hostIpAddress` | string | Yes |  |
| `prefixLength` | integer | Yes |  |
| `additionalHostIpSubnets` | array |  | Additional host IP subnets assigned to this VLAN. |
| `dhcpConfiguration` | Gateway Managed IPv4 DHCP Configuration |  | IPv4 DHCP configuration for this network. If this field is omitted or null, DHCP is not working and hosts must get an address statically or from another server in this broadcast domain. |
| `natOutboundIpAddressConfiguration` | array |  | List of NAT Outbound Configurations defining which IP addresses are used for NAT translation. This array must contain all WAN interfaces with `static` or `PPPoE` IPv4 connection configuration. |


### Gateway Managed IPv4 DHCP Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | Yes |  |


### Gateway Managed IPv4 DHCP Server Configuration


### Gateway managed network details


### Gateway managed network overview


### Guest access authorization request

Authorizes network access to a guest client. Client must be a guest.
This action cancels existing active authorization (if exists), creates a new one with new limits
and resets guest traffic counters.


### Guest access authorization response


### Guest access details


### Guest access overview


### Guest access unauthorization request

Unauthorizes network access and disconnects a guest client.


### Guest access unauthorization response


### Guest authorization details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `authorizedAt` | string | Yes | Timestamp when the guest has been authorized |
| `authorizationMethod` | string | Yes | Guest authorization method (API, Voucher etc) |
| `expiresAt` | string | Yes | Timestamp when the guest will get automatically unauthorized |
| `dataUsageLimitMBytes` | integer |  | (Optional) data usage limit in megabytes |
| `rxRateLimitKbps` | integer |  | (Optional) download rate limit in kilobits per second |
| `txRateLimitKbps` | integer |  | (Optional) upload rate limit in kilobits per second |
| `usage` | Guest authorization usage details |  |  |


### Guest authorization usage details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `durationSec` | integer | Yes |  |
| `rxBytes` | integer | Yes |  |
| `txBytes` | integer | Yes |  |
| `bytes` | integer | Yes |  |


### Hotspot voucher creation request

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `count` | integer |  | Number of vouchers to generate |
| `name` | string | Yes | Voucher note, duplicated across all generated vouchers |
| `authorizedGuestLimit` | integer |  | (Optional) limit for how many different guests can use the same voucher to authorize network access |
| `timeLimitMinutes` | integer | Yes | How long (in minutes) the voucher will provide access to the network since authorization of the first guest. Subsequently connected guests, if allowed, will share the same expiration time. |
| `dataUsageLimitMBytes` | integer |  | (Optional) data usage limit in megabytes |
| `rxRateLimitKbps` | integer |  | (Optional) download rate limit in kilobits per second |
| `txRateLimitKbps` | integer |  | (Optional) upload rate limit in kilobits per second |


### Hotspot voucher detail page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Hotspot voucher details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `createdAt` | string | Yes |  |
| `name` | string | Yes | Voucher note, may contain duplicate values across multiple vouchers |
| `code` | string | Yes | Secret code to active the voucher using the Hotspot portal |
| `authorizedGuestLimit` | integer |  | (Optional) limit for how many different guests can use the same voucher to authorize network access |
| `authorizedGuestCount` | integer | Yes | For how many guests the voucher has been used to authorize network access |
| `activatedAt` | string |  | (Optional) timestamp when the voucher has been activated (authorization time of the first guest) |
| `expiresAt` | string |  | (Optional) timestamp when the voucher will become expired. All guests using the voucher will be unauthorized from network access |
| `expired` | boolean | Yes | Whether the voucher has been expired and can no longer be used to authorize network access |
| `timeLimitMinutes` | integer | Yes | How long (in minutes) the voucher will provide access to the network since authorization of the first guest. Subsequently connected guests, if allowed, will share the same expiration time. |
| `dataUsageLimitMBytes` | integer |  | (Optional) data usage limit in megabytes |
| `rxRateLimitKbps` | integer |  | (Optional) download rate limit in kilobits per second |
| `txRateLimitKbps` | integer |  | (Optional) upload rate limit in kilobits per second |


### IP ACL rule endpoint

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IP Address selector


### IP address range

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `start` | string | Yes |  |
| `stop` | string | Yes |  |


### IP address range selector


### IP address selector

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IP matching

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IPv4 DHCP Relay Configuration


### IPv4 DHCP Server Configuration


### IPv4 matching

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IPv6 Client Address Assignment

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `dhcpConfiguration` | DHCP Configuration for IPv6 Network |  | IPv6 DHCP configuration for this network. At least one addressing method must be active: either enable SLAAC or provide DHCP configuration. If this field is null, SLAAC must be enabled. |
| `slaacEnabled` | boolean | Yes | Allows devices to obtain IPv6 addresses via SLAAC (Stateless Address Autoconfiguration) without DHCPv6. At least one addressing method must be active: either enable SLAAC or provide DHCP configuration. |


### IPv6 Static Configuration


### IPv6 matching

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IntegerType


### Integration blackout schedule configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `days` | array | Yes |  |


### IntegrationAclRuleDevicesFilterDto


### IntegrationAclRulePageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationDerivedSiteToSiteTunnelMetadata


### IntegrationDeviceAdoptionRequestDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `macAddress` | string | Yes |  |
| `ignoreDeviceLimit` | boolean | Yes |  |


### IntegrationDeviceTagPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationDnsARecordCreateUpdateDto


### IntegrationDnsARecordDto


### IntegrationDnsAaaaRecordCreateUpdateDto


### IntegrationDnsAaaaRecordDto


### IntegrationDnsCnameRecordCreateUpdateDto


### IntegrationDnsCnameRecordDto


### IntegrationDnsForwardDomainPolicyCreateUpdateDto


### IntegrationDnsForwardDomainPolicyDto


### IntegrationDnsMxRecordCreateUpdateDto


### IntegrationDnsMxRecordDto


### IntegrationDnsPolicyPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationDnsSrvRecordCreateUpdateDto


### IntegrationDnsSrvRecordDto


### IntegrationDnsTxtRecordCreateUpdateDto


### IntegrationDnsTxtRecordDto


### IntegrationFirewallPolicyActionAllowDto


### IntegrationFirewallPolicyActionBlockDto


### IntegrationFirewallPolicyActionRejectDto


### IntegrationFirewallPolicyDestinationApplicationCategoryFilterDto


### IntegrationFirewallPolicyDestinationApplicationFilterDto


### IntegrationFirewallPolicyDestinationDomainFilterDto


### IntegrationFirewallPolicyDestinationIpAddressFilterDto


### IntegrationFirewallPolicyDestinationIpv6IidFilterDto


### IntegrationFirewallPolicyDestinationNetworkFilterDto


### IntegrationFirewallPolicyDestinationPortFilterDto


### IntegrationFirewallPolicyDestinationRegionFilterDto


### IntegrationFirewallPolicyDestinationSiteToSiteVpnTunnelFilterDto


### IntegrationFirewallPolicyDestinationVpnServerFilterDto


### IntegrationFirewallPolicyIpAddressTrafficMatchingListFilterDto


### IntegrationFirewallPolicyIpMatchingIpAddressDto


### IntegrationFirewallPolicyIpMatchingRangeDto


### IntegrationFirewallPolicyIpMatchingSubnetDto


### IntegrationFirewallPolicyIpv4AndIpv6NamedProtocolDefaultDto


### IntegrationFirewallPolicyIpv4AndIpv6NamedProtocolFilterDto


### IntegrationFirewallPolicyIpv4AndIpv6ProtocolPresetFilterDto


### IntegrationFirewallPolicyIpv4AndIpv6ProtocolPresetTcpUdpDto


### IntegrationFirewallPolicyIpv4AndIpv6ProtocolScopeDto


### IntegrationFirewallPolicyIpv4NamedProtocolDefaultDto


### IntegrationFirewallPolicyIpv4NamedProtocolFilterDto


### IntegrationFirewallPolicyIpv4NamedProtocolIcmpDto


### IntegrationFirewallPolicyIpv4ProtocolPresetFilterDto


### IntegrationFirewallPolicyIpv4ProtocolPresetTcpUdpDto


### IntegrationFirewallPolicyIpv4ProtocolScopeDto


### IntegrationFirewallPolicyIpv6NamedProtocolDefaultDto


### IntegrationFirewallPolicyIpv6NamedProtocolFilterDto


### IntegrationFirewallPolicyIpv6NamedProtocolIcmpv6Dto


### IntegrationFirewallPolicyIpv6ProtocolPresetFilterDto


### IntegrationFirewallPolicyIpv6ProtocolPresetTcpUdpDto


### IntegrationFirewallPolicyIpv6ProtocolScopeDto


### IntegrationFirewallPolicyOrderingDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `orderedFirewallPolicyIds` | Ordered firewall policy IDs | Yes |  |


### IntegrationFirewallPolicyPortReferenceFilterDto


### IntegrationFirewallPolicyPortValueFilterDto


### IntegrationFirewallPolicySourceIpAddressFilterDto


### IntegrationFirewallPolicySourceIpv6IidFilterDto


### IntegrationFirewallPolicySourceMacAddressFilterDto


### IntegrationFirewallPolicySourceNetworkFilterDto


### IntegrationFirewallPolicySourcePortFilterDto


### IntegrationFirewallPolicySourceRegionFilterDto


### IntegrationFirewallPolicySourceSiteToSiteVpnTunnelFilterDto


### IntegrationFirewallPolicySourceVpnServerFilterDto


### IntegrationFirewallPolicySpecificDomainFilterDto


### IntegrationFirewallPolicySpecificIpAddressFilterDto


### IntegrationFirewallScheduleCustomDto


### IntegrationFirewallScheduleEveryDayDto


### IntegrationFirewallScheduleEveryWeekDto


### IntegrationFirewallScheduleOneTimeOnlyDto


### IntegrationGatewayManagedNetworkCreateUpdateDto


### IntegrationIotOptimizedWifiBroadcastCreateUpdateDto


### IntegrationIotOptimizedWifiBroadcastDetailDto


### IntegrationIotOptimizedWifiBroadcastOverviewDto


### IntegrationIpAclRuleCreateUpdateDto


### IntegrationIpAclRuleDto


### IntegrationIpAclRuleNetworkEndpointFilterDto


### IntegrationIpAclRulePortEndpointFilterDto


### IntegrationIpAclRuleSubnetEndpointFilterDto


### IntegrationIpV4TrafficMatchingListCreateUpdateDto


### IntegrationIpV4TrafficMatchingListDto


### IntegrationIpV6TrafficMatchingListCreateUpdateDto


### IntegrationIpV6TrafficMatchingListDto


### IntegrationIpv6AddressSuffixRangeSelectorDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `start` | string | Yes | Start suffix of the DHCPv6 address pool. |
| `stop` | string | Yes | End suffix of the DHCPv6 address pool. |


### IntegrationL2tpServerOverviewDto


### IntegrationLagMemberDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `deviceId` | string | Yes |  |
| `portIdxs` | array | Yes |  |


### IntegrationLagPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationLocalLagGlobalDto


### IntegrationLocalLagLocalDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `portIdxs` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### IntegrationMacAclRuleCreateUpdateDto


### IntegrationMacAclRuleDto


### IntegrationMacAclRuleMacAddressEndpointFilterDto


### IntegrationMcLagDomainDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `peers` | array | Yes |  |
| `lags` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### IntegrationMcLagDomainDtoPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationMcLagGlobalDto


### IntegrationMcLagLocalDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `members` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### IntegrationMcLagPeerDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `role` | string | Yes |  |
| `deviceId` | string | Yes |  |
| `linkPortIdxs` | array | Yes |  |


### IntegrationOpenVpnServerOverviewDto


### IntegrationPortTrafficMatchingListCreateUpdateDto


### IntegrationPortTrafficMatchingListDto


### IntegrationPptpServerOverviewDto


### IntegrationSiteToSiteIpsecTunnelOverviewDto


### IntegrationSiteToSiteOpenVpnTunnelOverviewDto


### IntegrationSiteToSiteVpnTunnelOverviewPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationSiteToSiteWireguardTunnelOverviewDto


### IntegrationStandardWifiBroadcastCreateUpdateDto


### IntegrationStandardWifiBroadcastDetailDto


### IntegrationStandardWifiBroadcastOverviewDto


### IntegrationSwitchManagedNetworkCreateUpdateDto


### IntegrationSwitchStackDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `members` | array | Yes |  |
| `lags` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### IntegrationSwitchStackDtoPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationSwitchStackLagGlobalDto


### IntegrationSwitchStackLagLocalDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `members` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### IntegrationSwitchStackMemberDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `deviceId` | string | Yes |  |


### IntegrationUidVpnServerOverviewDto


### IntegrationUnmanagedNetworkCreateUpdateDto


### IntegrationVoucherCreationResultDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `vouchers` | array |  |  |


### IntegrationVpnServerOverviewPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationWifiBasicDataRateConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `2.4` | integer | Yes |  |
| `5` | integer | Yes |  |


### IntegrationWifiBlackoutScheduleConfigurationPerAllDayDto


### IntegrationWifiBlackoutScheduleConfigurationPerDayWithTimeRangeDto


### IntegrationWifiBlackoutScheduleConfigurationTimeRangeDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `startTime` | string | Yes | Start time in 24-hour format (HH:mm) |
| `endTime` | string | Yes | End time in 24-hour format (HH:mm) |


### IntegrationWifiBroadcastPageDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### IntegrationWifiCaptivePortalConfigurationDetailDto


### IntegrationWifiClientFilteringPolicyDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |
| `macAddressFilter` | array | Yes |  |


### IntegrationWifiDerivedNasIdDto


### IntegrationWifiDeviceTagsFilterDto


### IntegrationWifiDevicesFilterDto


### IntegrationWifiDnsAssistanceAutoConfigurationDto


### IntegrationWifiDnsAssistanceManualConfigurationDto


### IntegrationWifiDtimPeriodConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `2.4` | integer | Yes |  |
| `5` | integer | Yes |  |
| `6` | integer | Yes |  |


### IntegrationWifiEnterpriseRadiusConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `profileId` | string | Yes |  |
| `nasId` | Wifi Radius NAS ID configuration | Yes |  |
| `macAuthenticationConfiguration` | IntegrationWifiRadiusMacAuthenticationConfigurationDto |  |  |


### IntegrationWifiHotspotConfigurationOverviewDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### IntegrationWifiMdnsProxyAllowPolicyDto


### IntegrationWifiMdnsProxyAutoConfigurationDto


### IntegrationWifiMdnsProxyBlockPolicyDto


### IntegrationWifiMdnsProxyCustomConfigurationDto


### IntegrationWifiMdnsProxyCustomServiceDto


### IntegrationWifiMdnsProxyPredefinedServiceDto


### IntegrationWifiMulticastFilteringAllowPolicyDto


### IntegrationWifiMulticastFilteringBlockPolicyDto


### IntegrationWifiNativeNetworkDto


### IntegrationWifiNonEnterpriseRadiusConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `profileId` | string | Yes |  |
| `nasId` | Wifi Radius NAS ID configuration | Yes |  |
| `macAuthenticationConfiguration` | IntegrationWifiRadiusMacAuthenticationConfigurationDto | Yes |  |


### IntegrationWifiOpenSecurityConfigurationDetailDto


### IntegrationWifiOpenSecurityConfigurationOverviewDto


### IntegrationWifiPasspointConfigurationDetailDto


### IntegrationWifiPresharedKeyDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `network` | Wifi network reference | Yes |  |
| `passphrase` | string | Yes |  |


### IntegrationWifiRadiusMacAuthenticationConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `macAddressFormat` | string | Yes |  |


### IntegrationWifiSaeConfigurationDto

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `anticloggingThresholdSeconds` | integer | Yes |  |
| `syncTimeSeconds` | integer | Yes |  |


### IntegrationWifiSpecificNetworkDto


### IntegrationWifiUserDefinedNasIdDto


### IntegrationWifiWpa2EnterpriseSecurityConfigurationDetailDto


### IntegrationWifiWpa2EnterpriseSecurityConfigurationOverviewDto


### IntegrationWifiWpa2PersonalSecurityConfigurationDetailDto


### IntegrationWifiWpa2PersonalSecurityConfigurationOverviewDto


### IntegrationWifiWpa2Wpa3EnterpriseSecurityConfigurationDetailDto


### IntegrationWifiWpa2Wpa3EnterpriseSecurityConfigurationOverviewDto


### IntegrationWifiWpa2Wpa3PersonalSecurityConfigurationDetailDto


### IntegrationWifiWpa2Wpa3PersonalSecurityConfigurationOverviewDto


### IntegrationWifiWpa3EnterpriseSecurityConfigurationDetailDto


### IntegrationWifiWpa3EnterpriseSecurityConfigurationOverviewDto


### IntegrationWifiWpa3PersonalSecurityConfigurationDetailDto


### IntegrationWifiWpa3PersonalSecurityConfigurationOverviewDto


### IntegrationWireguardServerOverviewDto


### LAG details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `members` | array | Yes |  |
| `metadata` | User defined entity metadata | Yes |  |


### Latest statistics for a device

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `uptimeSec` | integer |  |  |
| `lastHeartbeatAt` | string |  |  |
| `nextHeartbeatAt` | string |  |  |
| `loadAverage1Min` | number |  |  |
| `loadAverage5Min` | number |  |  |
| `loadAverage15Min` | number |  |  |
| `cpuUtilizationPct` | number |  |  |
| `memoryUtilizationPct` | number |  |  |
| `uplink` | Latest statistics for a device uplink interface |  |  |
| `interfaces` | Latest statistics for device interfaces | Yes |  |


### Latest statistics for a device uplink interface

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `txRateBps` | integer |  |  |
| `rxRateBps` | integer |  |  |


### Latest statistics for device interfaces

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `radios` | array |  |  |


### Latest statistics for wireless radio

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `frequencyGHz` | number | Yes |  |
| `txRetriesPct` | number |  |  |


### Local client access details

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Local client access overview

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### MAC ACL rule endpoint

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Multicast filtering policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |


### NAT Outbound Auto Configuration


### NAT Outbound Static Configuration


### Network DHCP Guarding

Details about DHCP Guarding settings for this Network.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `trustedDhcpServerIpAddresses` | array | Yes | List of trusted DHCP server IP addresses. |


### Network IPv6 Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interfaceType` | string | Yes |  |
| `clientAddressAssignment` | IPv6 Client Address Assignment | Yes | Client Address Assignment |
| `routerAdvertisement` | Router advertisement Configuration |  | Router advertisement. Without it, hosts will not autoconfigure addresses and will lack a default route even with DHCPv6. |
| `dnsServerIpAddressesOverride` | array |  | The IPv6 DNS server addresses assigned to this Network. If none are specified, they will be selected automatically. |
| `additionalHostIpSubnets` | array |  | Additional host IP subnets assigned to this VLAN. |


### Network details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `management` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `vlanId` | integer | Yes | VLAN ID. Must be 1 for the default network and >= 2 for additional networks. |
| `metadata` | User or system defined or orchestrated entity metadata | Yes | Orchestrated or System-defined configurable network support |
| `dhcpGuarding` | Network DHCP Guarding |  | DHCP Guarding settings for this Network. If this field is omitted or null, the feature is disabled |
| `default` | boolean | Yes |  |


### Network overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `management` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `vlanId` | integer | Yes | VLAN ID. Must be 1 for the default network and >= 2 for additional networks. |
| `metadata` | User or system defined or orchestrated entity metadata | Yes | Orchestrated or System-defined configurable network support |
| `default` | boolean | Yes |  |


### Network overview page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Network reference detail

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `referenceId` | string | Yes |  |


### Network reference resource

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `resourceType` | string | Yes |  |
| `referenceCount` | integer | Yes | Number of references of this type |
| `references` | array |  | List of references, present only if resourceType has API model defined |


### Network references

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `referenceResources` | array | Yes | List of network reference resources |


### NotFilterExpression


### Number port matching


### Number range port matching


### Orchestrated entity metadata


### Ordered firewall policy IDs

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `beforeSystemDefined` | array | Yes |  |
| `afterSystemDefined` | array | Yes |  |


### PXE Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `serverIpAddress` | string | Yes |  |
| `filename` | string | Yes |  |


### Patch firewall policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `loggingEnabled` | boolean |  |  |


### Port PoE overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `standard` | string | Yes |  |
| `type` | integer | Yes |  |
| `enabled` | boolean | Yes | Whether the PoE feature is enabled on the port |
| `state` | string | Yes | Whether the port currently supplies power to the (connected) device. |


### Port PoE power-cycle request


### Port action request

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |


### Port matching

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Port overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `idx` | integer | Yes |  |
| `state` | string | Yes |  |
| `connector` | string | Yes |  |
| `maxSpeedMbps` | integer | Yes |  |
| `speedMbps` | integer |  |  |
| `poe` | Port PoE overview |  |  |


### Prefix delegation IPv6 Configuration


### PropertyFilterExpression


### Radius Profile Overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `metadata` | User or system defined or derived entity metadata | Yes |  |


### Radius Profile Overview Page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Router advertisement Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `priority` | string | Yes | Router advertisement priority. |


### ScalarType


### SetType


### Site overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `internalReference` | string | Yes | Internal unique name of the site used in older APIs |
| `name` | string | Yes |  |


### Site overview page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Site-to-site VPN tunnel metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### Site-to-site VPN tunnel overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `metadata` | Site-to-site VPN tunnel metadata | Yes |  |


### StringType


### Subnet IPv4 matching


### Subnet IPv6 matching


### Switch Managed IPv4 Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `autoScaleEnabled` | boolean | Yes | Whether the Network can automatically scale its subnet size based on the number of active DHCP leases. |
| `hostIpAddress` | string | Yes |  |
| `prefixLength` | integer | Yes |  |
| `additionalHostIpSubnets` | array |  | Additional host IP subnets assigned to this VLAN. |
| `dhcpConfiguration` | Switch Managed IPv4 DHCP Configuration |  | IPv4 DHCP configuration for this network. If this field is omitted or null, DHCP is not working and hosts must get an address statically or from another server in this broadcast domain. |


### Switch Managed IPv4 DHCP Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | Yes |  |


### Switch managed network details


### Switch managed network overview


### Switching feature overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `lags` | array | Yes |  |


### System defined entity metadata


### Teleport client (connection) details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `connectedAt` | string |  |  |
| `ipAddress` | string |  |  |
| `access` | Teleport client access details | Yes |  |


### Teleport client (connection) overview


### Teleport client access details

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Teleport client access overview

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### TimestampType


### Traffic matching list

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |


### Traffic matching lists page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### UUIDType


### Unmanaged network details


### Unmanaged network overview


### User defined entity metadata


### User defined or derived entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### User or derived or orchestrated entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### User or orchestrated entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### User or system defined entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### User or system defined or derived entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### User or system defined or orchestrated entity metadata

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `origin` | string | Yes |  |


### VPN client (connection) details


### VPN client (connection) overview


### VPN client access details

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### VPN client access overview

Represents the type of network access and/or any applicable authorization status the client is using.

- **Wired clients** may have direct access without additional authorization.
- **Wireless clients** can be connected via a protected network or an open network
  that may require additional authorization (e.g., a guest portal).
- **VPN clients** may have different authorization mechanisms.

Currently, the only two supported access types are `GUEST` (used for wired and wireless guest clients)
and `DEFAULT` (a placeholder, which might be refined in the future releases, used for all other clients).

Filtering is possible by `access.type`, for example `access.type.eq('GUEST')` to list guest clients.

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### VPN server overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `metadata` | User defined or derived entity metadata | Yes |  |


### Voucher deletion results

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `vouchersDeleted` | integer |  |  |


### WAN NAT Outbound Configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `wanInterfaceId` | string | Yes |  |


### WAN overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | Yes |  |
| `name` | string | Yes |  |


### WAN overview page

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `offset` | integer | Yes |  |
| `limit` | integer | Yes |  |
| `count` | integer | Yes |  |
| `totalCount` | integer | Yes |  |
| `data` | array | Yes |  |


### Wifi Radius NAS ID configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Wifi broadcast create or update

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `name` | string | Yes |  |
| `network` | Wifi network reference |  |  |
| `enabled` | boolean | Yes |  |
| `securityConfiguration` | Wifi security configuration detailObject | Yes |  |
| `broadcastingDeviceFilter` | Broadcasting device filter |  | Defines the custom scope of devices that will broadcast this WiFi network. If null, the WiFi network will be broadcast by all Access Point capable devices. |
| `mdnsProxyConfiguration` | mDNS filtering configuration |  |  |
| `multicastFilteringPolicy` | Multicast filtering policy |  |  |
| `multicastToUnicastConversionEnabled` | boolean | Yes |  |
| `clientIsolationEnabled` | boolean | Yes |  |
| `hideName` | boolean | Yes |  |
| `uapsdEnabled` | boolean | Yes | Indicates whether Unscheduled Automatic Power Save Delivery (U-APSD) is enabled |
| `basicDataRateKbpsByFrequencyGHz` | IntegrationWifiBasicDataRateConfigurationDto |  |  |
| `clientFilteringPolicy` | IntegrationWifiClientFilteringPolicyDto |  | Client connection filtering policy. Allow/restrict access to the WiFi network based on client device MAC addresses. |
| `blackoutScheduleConfiguration` | Integration blackout schedule configuration |  |  |


### Wifi broadcast details

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `metadata` | User or derived or orchestrated entity metadata | Yes |  |
| `enabled` | boolean | Yes |  |
| `network` | Wifi network reference |  |  |
| `securityConfiguration` | Wifi security configuration detailObject | Yes |  |
| `broadcastingDeviceFilter` | Broadcasting device filter |  | Defines the custom scope of devices that will broadcast this WiFi network. If null, the WiFi network will be broadcast by all Access Point capable devices. |
| `mdnsProxyConfiguration` | mDNS filtering configuration |  |  |
| `multicastFilteringPolicy` | Multicast filtering policy |  |  |
| `multicastToUnicastConversionEnabled` | boolean | Yes |  |
| `clientIsolationEnabled` | boolean | Yes |  |
| `hideName` | boolean | Yes |  |
| `uapsdEnabled` | boolean | Yes | Indicates whether Unscheduled Automatic Power Save Delivery (U-APSD) is enabled |
| `basicDataRateKbpsByFrequencyGHz` | IntegrationWifiBasicDataRateConfigurationDto |  |  |
| `clientFilteringPolicy` | IntegrationWifiClientFilteringPolicyDto |  | Client connection filtering policy. Allow/restrict access to the WiFi network based on client device MAC addresses. |
| `blackoutScheduleConfiguration` | Integration blackout schedule configuration |  |  |


### Wifi broadcast overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `id` | string | Yes |  |
| `name` | string | Yes |  |
| `enabled` | boolean | Yes |  |
| `metadata` | User or derived or orchestrated entity metadata | Yes |  |
| `network` | Wifi network reference |  |  |
| `securityConfiguration` | Wifi security configuration overview | Yes |  |
| `broadcastingDeviceFilter` | Broadcasting device filter |  | Defines the custom scope of devices that will broadcast this WiFi network. If null, the WiFi network will be broadcast by all Access Point capable devices. |


### Wifi hotspot configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Wifi network reference

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Wifi security configuration detailObject

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `radiusConfiguration` |  |  |  |


### Wifi security configuration overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |


### Wired client details


### Wired client overview


### Wireless client details


### Wireless client overview


### Wireless radio overview

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `wlanStandard` | string | Yes |  |
| `frequencyGHz` | number | Yes |  |
| `channelWidthMHz` | integer | Yes |  |
| `channel` | integer |  |  |


### mDNS filtering configuration

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | Yes |  |


### mDNS proxy policy

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | string | Yes |  |
| `deviceFilter` | Broadcasting device filter |  | Defines the custom scope of devices that will filter Mdns. If null, the mDNS filtering will be added to all Access Point capable devices. |


### mDNS service

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |

