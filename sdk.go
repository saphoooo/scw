package scw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// functions
// Order create, list, get, update, delete
func (c *Config) CreateFlexibleIP(newFlexibleIP FlexibleIP) FlexibleIPResp {
	json_data, err := json.Marshal(newFlexibleIP)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
	req, err := http.NewRequest("POST", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/ips", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ipResp FlexibleIPResp
	err = json.Unmarshal(bodyBytes, &ipResp)
	if err != nil {
		log.Fatalln(err)
	}
	return ipResp
}

func (c *Config) CreateSrv(server *Server) SrvResp {
	json_data, err := json.Marshal(server)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var serverResponse SrvResp
	err = json.Unmarshal(bodyBytes, &serverResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return serverResponse
}

//func (c *Config) CreateSecurityGroup(newSecurityGroup SecurityGroup) SecurityGroupResp {
func (c *Config) CreateSecurityGroup(newSecurityGroup SecurityGroup) {
	json_data, err := json.Marshal(newSecurityGroup)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
	req, err := http.NewRequest("POST", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/security_groups", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	/*
		var sgResp SecurityGroupResp
		err = json.Unmarshal(bodyBytes, &sgResp)
		if err != nil {
			log.Fatalln(err)
		}
		return sgResp
	*/
}

func (c *Config) CreateSecurityGroupRule(SecurityGroupID string, newSecurityGroupRule SecurityGroupRule) SecurityGroupRuleResp {
	json_data, err := json.Marshal(newSecurityGroupRule)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
	req, err := http.NewRequest("POST", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/security_groups/"+SecurityGroupID+"/rules", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ruleResp SecurityGroupRuleResp
	err = json.Unmarshal(bodyBytes, &ruleResp)
	if err != nil {
		log.Fatalln(err)
	}
	return ruleResp
}

func (c *Config) AddUserData(serverID, userData string) {
	req, err := http.NewRequest("PATCH", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers/"+serverID+"/user_data/cloud-init", bytes.NewBuffer([]byte(userData)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) ListServers() {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) ListSecurityGroups() {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/security_groups", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) ListUserData(serverID string) {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers/"+serverID+"/user_data/cloud-init", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) ListSercurityGroupRules(sercurityGroupID string) {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/security_groups/"+sercurityGroupID+"/rules", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) GetSecurityGroup(sercurityGroupID string) {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/security_groups/"+sercurityGroupID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	/*
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var securityGroupResp SecurityGroupResp
		err = json.Unmarshal(bodyBytes, &securityGroupResp)
		if err != nil {
			log.Fatalln(err)
		}
		return securityGroupResp
	*/
}

func (c *Config) GetSrv(serverID string) SrvResp {
	req, err := http.NewRequest("GET", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers/"+serverID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var serverResp SrvResp
	err = json.Unmarshal(bodyBytes, &serverResp)
	if err != nil {
		log.Fatalln(err)
	}
	return serverResp
}

func (c *Config) UpdateSrvState(serverID, state string) {
	var newAction SrvAction
	newAction.Action = state
	json_data, err := json.Marshal(newAction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
	req, err := http.NewRequest("POST", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers/"+serverID+"/action", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) DeleteFlexibleIP(flexibleIPID string) {
	req, err := http.NewRequest("DELETE", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/ips/"+flexibleIPID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) DeleteVolume(volumeID string) {
	req, err := http.NewRequest("DELETE", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/volumes/"+volumeID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (c *Config) DeleteSrv(serverID string) {
	req, err := http.NewRequest("DELETE", "https://api.scaleway.com/instance/v1/zones/"+c.Zone+"/servers/"+serverID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

// structs
type Server struct {
	Name              string   `json:"name"`                          // The server name
	DynamicIPRequired bool     `json:"dynamic_ip_required,omitempty"` // Define if a dynamic IP is required for the instance
	CommercialType    string   `json:"commercial_type"`               // Define the server commercial type (i.e. GP1-S)
	Image             string   `json:"image,omitempty"`               // The server image ID
	Volumes           Volumes  `json:"volumes,omitempty"`             // The volumes attached to the server
	EnableIPv6        bool     `json:"enable_ipv6"`                   // True if IPv6 is enabled on the server
	PublicIP          string   `json:"public_ip,omitempty"`           // The ID of the reserved IP to attach to the server
	BootType          string   `json:"boot_type,omitempty"`           // The boot type to use. Possible values are local, bootscript and rescue. The default value is local
	Bootscript        string   `json:"bootscript,omitempty"`          // The bootscript ID to use when `boot_type` is set to `bootscript
	Organization      string   `json:"organization,omitempty"`        // The server organization ID. Only one of organization and project may be set, organisation is deprecated
	Project           string   `json:"project,omitempty"`             // The server project ID. Only one of organization and project may be set
	Tags              []string `json:"tags,omitempty"`                // The server tags
	SecurityGroup     string   `json:"security_group,omitempty"`      // The security group ID
	PlacementGroup    string   `json:"placement_group,omitempty"`     // Placement group ID if server must be part of a placement group
}

type Volumes struct {
	Zero *Zero `json:"0,omitempty"`
}

type Zero struct {
	ID           string `json:"id"`                     // UUID of the volume
	Name         string `json:"name"`                   // Name of the volume
	Size         int    `json:"size"`                   // Disk size of the volume (in bytes)
	VolumeType   string `json:"volume_type"`            // Type of the volume. Possible values are l_ssd and b_ssd. The default value is l_ssd
	Organization string `json:"organization,omitempty"` // Organization ID of the volume. Only one of organization and project may be set
	Project      string `json:"project,omitempty"`      // Project ID of the volume. Only one of organization and project may be set
}

/*
type response struct {
	Server serverResponse `json:"server,omitempty"`
}

type serverResponse struct {
	ID                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Organization      string   `json:"organization,omitempty"`
	Project           string   `json:"project,omitempty"`
	AllowedActions    []string `json:"allowed_actions,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	CommercialType    string   `json:"commercial_type,omitempty"`
	CreationDate      string   `json:"creation_date,omitempty"`
	DynamicIPRequired bool     `json:"dynamic_ip_required,omitempty"`
	EnableIPv6        bool     `json:"enable_ipv6,omitempty"`
	Hostname          string   `json:"hostname,omitempty"`
	Image             image    `json:"image,omitempty"`
	Protected         bool     `json:"protected,omitempty"`
	PrivateIP         string   `json:"private_ip,omitempty"`
	PublicIP          struct {
		ID      string `json:"id,omitempty"`
		Address string `json:"address,omitempty"`
		Dynamic bool   `json:"dynamic,omitempty"`
	} `json:"public_ip,omitempty"`
	ModificationDate string `json:"modification_date,omitempty"`
	State            string `json:"state,omitempty"`
	Location         struct {
		ClusterID    string `json:"cluster_id,omitempty"`
		HypervisorID string `json:"hypervisor_id,omitempty"`
		NodeID       string `json:"node_id,omitempty"`
		PlatformID   string `json:"platform_id,omitempty"`
		ZoneID       string `json:"zone_id,omitempty"`
	} `json:"location,omitempty"`
	IPv6 struct {
		Address string `json:"address,omitempty"`
		Gateway string `json:"gateway,omitempty"`
		Netmask string `json:"netmask,omitempty"`
	} `json:"ipv6,omitempty"`
	Bootscript struct {
		Bootcmdargs  string `json:"bootcmdargs,omitempty"`
		Default      bool   `json:"default,omitempty"`
		DTB          string `json:"dtb,omitempty"`
		ID           string `json:"id,omitempty"`
		Initrd       string `json:"initrd,omitempty"`
		Kernel       string `json:"kernel,omitempty"`
		Organization string `json:"organization,omitempty"`
		Project      string `json:"project,omitempty"`
		Public       bool   `json:"public,omitempty"`
		Title        string `json:"title,omitempty"`
		Arch         string `json:"arch,omitempty"`
		Zone         string `json:"zone,omitempty"`
	} `json:"bootscript,omitempty"`
	BootType      string       `json:"boot_type,omitempty"`
	Volumes       extraVolumes `json:"volumes,omitempty"`
	SecurityGroup struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"security_group,omitempty"`
	Maintenances   []string `json:"maintenances,omitempty"`
	StateDetail    string   `json:"state_detail,omitempty"`
	Arch           string   `json:"arch,omitempty"`
	PlacementGroup struct {
		ID              string `json:"id,omitempty"`
		Name            string `json:"name,omitempty"`
		Organization    string `json:"organization,omitempty"`
		Project         string `json:"project,omitempty"`
		PolicyMode      string `json:"policy_mode,omitempty"`
		PolicyType      string `json:"policy_type,omitempty"`
		PolicyTespected bool   `json:"policy_respected,omitempty"`
		Zone            string `json:"zone,omitempty"`
	} `json:"placement_group,omitempty"`
	PrivateNics []nics `json:"private_nics,omitempty"`
	Zone        string `json:"zone,omitempty"`
}

type nics struct {
	ID               string `json:"id,omitempty"`
	ServerID         string `json:"server_id,omitempty"`
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	MacAddress       string `json:"mac_address,omitempty"`
}

type image struct {
	ID                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Arch              string            `json:"arch,omitempty"`
	CreationDate      string            `json:"creation_date,omitempty"`
	ModificationDate  string            `json:"modification_date,omitempty"`
	DefaultBootscript defaultBootscript `json:"default_bootscript,omitempty"`
	ExtraVolumes      extraVolumes      `json:"extra_volumes,omitempty"`
	FromServer        string            `json:"from_server,omitempty"`
	Organization      string            `json:"organization,omitempty"`
	Public            bool              `json:"public,omitempty"`
	RootVolume        struct {
		ID         string `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
		Size       int    `json:"size,omitempty"`
		VolumeType string `json:"volume_type,omitempty"`
	} `json:"root_volume,omitempty"`
	State   string `json:"state,omitempty"`
	Project string `json:"project,omitempty"`
	Zone    string `json:"zone,omitempty"`
}

type defaultBootscript struct {
	Bootcmdargs  string `json:"bootcmdargs,omitempty"`
	Default      bool   `json:"default,omitempty"`
	DTB          string `json:"dtb,omitempty"`
	ID           string `json:"id,omitempty"`
	Initrd       string `json:"initrd,omitempty"`
	Kernel       string `json:"kernel,omitempty"`
	Organization string `json:"organization,omitempty"`
	Project      string `json:"project,omitempty"`
	Public       bool   `json:"public,omitempty"`
	Title        string `json:"title,omitempty"`
	Arch         string `json:"arch,omitempty"`
	Zone         string `json:"zone,omitempty"`
}

type extraVolumes struct {
	ExtraVolumeKey struct {
		ID               string `json:"id,omitempty"`
		Name             string `json:"name,omitempty"`
		ExportURI        string `json:"export_uri,omitempty"`
		Size             int    `json:"size,omitempty"`
		VolumeType       string `json:"volume_type,omitempty"`
		CreationDate     string `json:"creation_date,omitempty"`
		ModificationDate string `json:"modification_date,omitempty"`
		Organization     string `json:"organization,omitempty"`
		Project          string `json:"project,omitempty"`
		Server           struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"server,omitempty"`
		State string `json:"state,omitempty"`
		Zone  string `json:"zone,omitempty"`
	} `json:"extra_volumeKey,omitempty"`
}
*/

type SrvAction struct {
	Action string `json:"action"`         // poweron, backup, stop_in_place, poweroff, terminate and reboot
	Name   string `json:"name,omitempty"` // name of the backup
}

type FlexibleIP struct {
	Organization string   `json:"organization,omitempty"`
	Project      string   `json:"project"`
	Server       string   `json:"server,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

type FlexibleIPResp struct {
	IP       IP     `json:"ip"`
	Location string `json:"Location"`
}

type IP struct {
	ID      string `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Reverse string `json:"reverse,omitempty"`
	Server  struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"server,omitempty"`
	Organization string   ` json:"organization,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Project      string   ` json:"project,omitempty"`
	Zone         string   ` json:"zone,omitempty"`
}

type SrvResp struct {
	Server struct {
		ID                string   `json:"id"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		State             string   `json:"state,omitempty"`
		AllowedActions    []string `json:"allowed_actions,omitempty"`
		Tags              []string `json:"tags,omitempty"`
		CreationDate      string   `json:"creation_date,omitempty"`
		DynamicIPRequired bool     `json:"dynamic_ip_required,omitempty"`
		PrivateIP         string   `json:"private_ip,omitempty"`
		PublicIP          struct {
			ID      string `json:"id,omitempty"`
			Address string `json:"address,omitempty"`
			Dynamic bool   `json:"dynamic,omitempty"`
		} `json:"public_ip,omitempty"`
		Volumes struct {
			VolumeKey struct {
				ID string `json:"id"`
			} `json:"0"`
		} `json:"volumes"`
	} `json:"server"`
}

type SecurityGroup struct {
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	Organization          string `json:"organization,omitempty"`
	Project               string `json:"project,omitempty"`
	OrganizationDefault   bool   `json:"organization_default,omitempty"`
	ProjectDefault        bool   `json:"project_default,omitempty"`
	Stateful              bool   `json:"stateful,omitempty"`
	InboundDefaultPolicy  string `json:"inbound_default_policy,omitempty"`
	OutboundDefaultPolicy string `json:"outbound_default_policy,omitempty"`
	EnableDefaultSecurity bool   `json:"enable_default_security,omitempty"`
}

type SecurityGroupResp struct {
	SGroup struct {
		ID                    string `json:"id"`
		Name                  string `json:"name"`
		Description           string `json:"description,omitempty"`
		EnableDefaultSecurity bool   `json:"enable_default_security,omitempty"`
		InboundDefaultPolicy  string `json:"inbound_default_policy,omitempty"`
		OutboundDefaultPolicy string `json:"outbound_default_policy,omitempty"`
		Organization          string `json:"organization,omitempty"`
		Project               string `json:"project,omitempty"`
		OrganizationDefault   bool   `json:"organization_default,omitempty"`
		ProjectDefault        bool   `json:"project_default,omitempty"`
		CreationDate          string `json:"creation_date,omitempty"`
		ModificationDate      string `json:"modification_date,omitempty"`
		Servers               []struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"servers,omitempty"`
		Stateful bool   `json:"stateful,omitempty"`
		Zone     string `json:"zone,omitempty"`
	} `json:"security_group"`
}

type SecurityGroupRule struct {
	ID           string `json:"id,omitempty"`
	Protocol     string `json:"protocol"`
	Direction    string `json:"direction"`
	Action       string `json:"action"`
	IPRange      string `json:"ip_range"`
	DestPortFrom int    `json:"dest_port_from"`
	DestPortTo   int    `json:"dest_port_to"`
	Position     int    `json:"position"`
	Editable     bool   `json:"editable"`
	Zone         string `json:"zone,omitempty"`
}

type SecurityGroupRuleResp struct {
	Rule SecurityGroupRule `json:"rule"`
}

type Config struct {
	Token string `json:"token"`
	Zone  string `json:"zone"`
}
