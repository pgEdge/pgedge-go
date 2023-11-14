package dbspec

import (
	"encoding/json"
	"os"

	"github.com/pgEdge/pgedge-go/world"
)

type UserType string

const (
	InternalAdmin    UserType = "internal_admin"
	InternalReadOnly UserType = "internal_read_only"
	Admin            UserType = "admin"
	App              UserType = "app"
	AppReadOnly      UserType = "app_read_only"
	PoolerAuth       UserType = "pooler_auth"
	Other            UserType = "other"
)

type Node struct {
	Name             string         `json:"name"`
	Domain           string         `json:"domain"`
	Region           string         `json:"region"`
	InternalHostname string         `json:"internal_hostname"`
	Location         world.Location `json:"location"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Superuser bool     `json:"superuser"`
	Service   string   `json:"service,omitempty"`
	Type      UserType `json:"type,omitempty"`
}

type Database struct {
	ID               string  `json:"id"`
	TenantID         string  `json:"tenant_id"`
	ClusterID        string  `json:"cluster_id"`
	FriendlyName     string  `json:"friendly_name"`
	Name             string  `json:"name"`
	Port             int     `json:"port"`
	DomainPrefix     string  `json:"domain_prefix"`
	ReservedCapacity int     `json:"reserved_capacity"`
	Nodes            []*Node `json:"nodes"`
	Users            []*User `json:"users"`
	Self             *Node   `json:"self,omitempty"`
}

func (db *Database) GetLogin(userType UserType) (*Login, bool) {
	for _, user := range db.Users {
		if user.Type == userType && user.Service == "postgres" {
			return &Login{Username: user.Username, Password: user.Password}, true
		}
	}
	return nil, false
}

func ReadSpec(path string) (*Database, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec Database
	if err = json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
