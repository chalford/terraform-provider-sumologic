package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider My first provider!
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dummy_server": resourceServer(),
		},
	}
}
