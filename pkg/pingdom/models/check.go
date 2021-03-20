package models

type Check struct {
	ID                  int    `json:"id"`
	Created             int    `json:"created"`
	Name                string `json:"name"`
	Hostname            string `json:"hostname"`
	ResolutionInMinutes int    `json:"resolution"`
	Type                string `json:"type"`
	Ipv6                bool   `json:"ipv6"`
	VerifyCertificate   bool   `json:"verify_certificate"`
	Tags                []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"tags"`
	Maintenances []string `json:"maintenanceids"`
}
