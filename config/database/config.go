package database

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nuodbaas_database", func(r *config.Resource) {
		// Get rid of any attributes that cannot be processed
		for key, attribute := range r.TerraformResource.Schema {
			if attribute.Type == schema.TypeInvalid {
				delete(r.TerraformResource.Schema, key)
			}
		}
	})
}
