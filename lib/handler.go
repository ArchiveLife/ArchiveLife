package lib

import (
	"encoding/json"
	"net/http"

	"github.com/ArchiveLife/core/adapter"
)

type serviceDto struct {
	Name        string
	Description string
	Options     []*adapter.Option
}

type serviceOption struct {
	Name string
}

type JSONObject map[string]interface{}

func APIGetServices(w http.ResponseWriter, r *http.Request) {
	services := []serviceDto{}

	for _, s := range GetServices() {
		services = append(services, serviceDto{
			Name:        s.GetName(),
			Description: s.GetDescription(),
			Options:     s.GetOptions(),
		})
	}

	json.NewEncoder(w).Encode(JSONObject{
		"services": services,
	})
}
