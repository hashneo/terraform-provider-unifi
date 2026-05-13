package client

import "errors"

// ── Protect API types (per protect integration spec v7.0.107) ─────────────────
//
// The cloud proxy returns bare JSON arrays (not {"data":[...]}).
// The local controller wraps in {"data":[...]}.
// doProtectList handles both.

// Camera maps to the "camera" schema.
type Camera struct {
	ID              string `json:"id"`
	ModelKey        string `json:"modelKey"`
	Name            string `json:"name"`
	MAC             string `json:"mac"`
	State           string `json:"state"`
	IsMicEnabled    bool   `json:"isMicEnabled"`
}

// NVR maps to the "nvr" schema.
type NVR struct {
	ID       string `json:"id"`
	ModelKey string `json:"modelKey"`
	Name     string `json:"name"`
}

// Sensor maps to the "sensor" schema.
type Sensor struct {
	ID         string         `json:"id"`
	ModelKey   string         `json:"modelKey"`
	Name       string         `json:"name"`
	MAC        string         `json:"mac"`
	State      string         `json:"state"`
	IsOpened   *bool          `json:"isOpened"`
	IsMotion   bool           `json:"isMotionDetected"`
	Battery    *BatteryStatus `json:"batteryStatus"`
}

// BatteryStatus maps to the "batteryStatus" schema.
type BatteryStatus struct {
	Percentage int  `json:"percentage"`
	IsLow      bool `json:"isLow"`
	IsCharging bool `json:"isCharging"`
}

// Light maps to the "light" schema.
type Light struct {
	ID               string `json:"id"`
	ModelKey         string `json:"modelKey"`
	Name             string `json:"name"`
	MAC              string `json:"mac"`
	State            string `json:"state"`
	IsLightOn        bool   `json:"isLightOn"`
	IsPirMotion      bool   `json:"isPirMotionDetected"`
}

// Viewer maps to the "viewer" schema.
type Viewer struct {
	ID       string `json:"id"`
	ModelKey string `json:"modelKey"`
	Name     string `json:"name"`
	MAC      string `json:"mac"`
	State    string `json:"state"`
}

// ── Protect API methods ───────────────────────────────────────────────────────

// doProtectList fetches a protect endpoint that returns either a bare array
// or a {"data":[...]} envelope depending on whether using cloud proxy or local.
func doProtectList[T any](c *Client, path string) ([]T, error) {
	// Try bare array first (cloud proxy).
	var items []T
	if err := c.doProtect(path, &items); err == nil {
		return items, nil
	}
	// Fall back to wrapped envelope (local controller).
	var wrapped struct {
		Data []T `json:"data"`
	}
	if err := c.doProtect(path, &wrapped); err != nil {
		return nil, err
	}
	return wrapped.Data, nil
}

func (c *Client) ListCameras() ([]Camera, error) {
	return doProtectList[Camera](c, "/cameras")
}

// GetNVR returns the first NVR from /v1/nvrs (plural per spec).
func (c *Client) GetNVR() (*NVR, error) {
	nvrs, err := doProtectList[NVR](c, "/nvrs")
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if len(nvrs) == 0 {
		return nil, nil
	}
	return &nvrs[0], nil
}

func (c *Client) ListSensors() ([]Sensor, error) {
	return doProtectList[Sensor](c, "/sensors")
}

func (c *Client) ListLights() ([]Light, error) {
	return doProtectList[Light](c, "/lights")
}

func (c *Client) ListViewers() ([]Viewer, error) {
	return doProtectList[Viewer](c, "/viewers")
}
