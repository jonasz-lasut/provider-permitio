package permitio

import "github.com/crossplane/upjet/pkg/config"

type TypeMap = map[string]interface{}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("permitio_condition_set_rule", func(r *config.Resource) {
		r.References["resource_set"] = config.Reference{
			TerraformName: "permitio_resource_set",
		}
		r.References["user_set"] = config.Reference{
			TerraformName: "permitio_user_set",
		}
		r.ShortGroup = "permitio"
		r.Kind = "ConditionSetRule"
	})

	p.AddResourceConfigurator("permitio_resource_set", func(r *config.Resource) {
		r.ShortGroup = "permitio"
		r.Kind = "ResourceSet"
	})

	p.AddResourceConfigurator("permitio_relation", func(r *config.Resource) {
		r.References["object_resource"] = config.Reference{
			TerraformName: "permitio_resource",
		}
		r.References["subject_resource"] = config.Reference{
			TerraformName: "permitio_resource",
		}
	})

	p.AddResourceConfigurator("permitio_role_derivation", func(r *config.Resource) {
		r.References["linked_by"] = config.Reference{
			TerraformName: "permitio_relation",
		}
		r.References["on_resource"] = config.Reference{
			TerraformName: "permitio_resource",
		}
		r.References["resource"] = config.Reference{
			TerraformName: "permitio_resource",
		}
		r.References["role"] = config.Reference{
			TerraformName: "permitio_role",
		}
		r.References["to_role"] = config.Reference{
			TerraformName: "permitio_role",
		}
		r.ShortGroup = "permitio"
		r.Kind = "RoleDerivation"
	})

	p.AddResourceConfigurator("permitio_user_set", func(r *config.Resource) {
		r.ShortGroup = "permitio"
		r.Kind = "UserSet"
	})

	p.AddResourceConfigurator("permitio_proxy_config", func(r *config.Resource) {
		r.ShortGroup = "permitio"
		r.Kind = "ProxyConfig"
	})
}
