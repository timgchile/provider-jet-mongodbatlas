/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mongodba

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/timgchile/provider-jet-mongodba/config/common"
)

// Configure configures the root group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_cluster", func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}

			if k == "project_id" {
				ref := config.Reference{
					Type:      common.APISPackagePath + "/mongodbatlas/v1alpha1.Project",
					Extractor: common.ExtractResourceIDFuncPath,
				}
				if r.ShortGroup == "" {
					ref.Type = "Project"
				}
				r.References["project_id"] = ref
			}
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("mongodbatlas_advanced_cluster", func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}

			if k == "project_id" {
				ref := config.Reference{
					Type:      common.APISPackagePath + "/mongodbatlas/v1alpha1.Project",
					Extractor: common.ExtractResourceIDFuncPath,
				}
				if r.ShortGroup == "" {
					ref.Type = "Project"
				}
				r.References["project_id"] = ref
			}
		}

		r.UseAsync = true
	})
}
