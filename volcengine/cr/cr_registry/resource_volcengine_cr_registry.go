package cr_registry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR Instance can be imported using the name, e.g.
```
$ terraform import volcengine_cr_instance.default enterprise-x
```

*/

func ResourceVolcengineCrRegistry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrRegistryCreate,
		Read:   resourceVolcengineCrRegistryRead,
		Update: resourceVolcengineCrRegistryUpdate,
		Delete: resourceVolcengineCrRegistryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of registry.",
			},
			"delete_immediately": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether delete registry immediately.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of registry user.",
			},
		},
	}
	dataSource := DataSourceVolcengineCrRegistries().Schema["registries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrRegistryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on creating CrRegistry %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrRegistryRead(d, meta)
}

func resourceVolcengineCrRegistryUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on updating CrRegistry  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrRegistryRead(d, meta)
}

func resourceVolcengineCrRegistryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on deleting CrRegistry %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrRegistryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("Error on reading CrRegistry %q,%s", d.Id(), err)
	}
	return err
}
