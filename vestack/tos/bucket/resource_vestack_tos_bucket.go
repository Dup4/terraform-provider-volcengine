package bucket

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VPC can be imported using the id, e.g.
```
$ terraform import vestack_vpc.default vpc-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackTosBucket() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackTosBucketCreate,
		Read:   resourceVestackTosBucketRead,
		Update: resourceVestackTosBucketUpdate,
		Delete: resourceVestackTosBucketDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the bucket.",
			},
			"public_acl": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"private",
					"public-read",
					"public-read-write",
					"authenticated-read",
					"bucket-owner-read",
				}, false),
				Default:     "private",
				Description: "The public acl control of bucket.",
			},
			"storage_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"STANDARD",
					"IA",
				}, false),
				Default:     "STANDARD",
				Description: "The storage type of the bucket.",
			},
			"enable_version": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The flag of enable tos version.",
			},
			"account_acl": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The user set of grant full control.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"acl_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "CanonicalUser",
							ValidateFunc: validation.StringInSlice([]string{
								"CanonicalUser",
							}, false),
						},
						"permission": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"FULL_CONTROL",
								"READ",
								"READ_ACP",
								"WRITE",
								"WRITE_ACP",
							}, false),
						},
					},
				},
				Set: tosAccountAclHash,
			},
		},
	}
	return resource
}

func resourceVestackTosBucketCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Create(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on creating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosBucketRead(d, meta)
}

func resourceVestackTosBucketRead(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Read(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on reading tos bucket %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackTosBucketUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Update(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on updating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosBucketRead(d, meta)
}

func resourceVestackTosBucketDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Delete(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on deleting tos bucket %q, %s", d.Id(), err)
	}
	return err
}
