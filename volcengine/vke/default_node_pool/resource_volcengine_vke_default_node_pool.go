package default_node_pool

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func ResourceVolcengineDefaultNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineDefaultNodePoolCreate,
		Update: resourceVolcengineDefaultNodePoolUpdate,
		Read:   resourceVolcengineDefaultNodePoolUpdate,
		Delete: resourceVolcengineNodePoolDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ClusterId of NodePool.",
			},
			"instances": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      defaultNodePoolNodeHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The instance id.",
						},
						"keep_instance_name": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "The flag of keep instance name, the value is `true` or `false`.Default is `false`.",
						},
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The Image Id to the ECS Instance.",
						},
						"container_storage_path": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: defaultNodePoolDiffSuppress(),
							Description:      "The container storage path.When additional_container_storage_enabled is `false` will ignore.",
						},
					},
				},
				Description: "The ECS InstanceIds add to NodePool.",
			},
			"kubernetes_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Key of Labels.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Value of Labels.",
									},
								},
							},
							Description: "The Labels of KubernetesConfig.",
						},
						"taints": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Key of Taints.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Value of Taints.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Effect of Taints.",
									},
								},
							},
							Description: "The Taints of KubernetesConfig.",
						},
						"cordon": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "The Cordon of KubernetesConfig.",
						},
					},
				},
				Description: "The KubernetesConfig of NodeConfig.",
			},
			"node_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The SecurityGroupIds of Security.",
									},
									"security_strategies": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The SecurityStrategies of Security.",
									},
									"login": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Password of Security.",
												},
												"ssh_key_pair_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The SshKeyPairName of Security.",
												},
											},
										},
										Description: "The Login of Security.",
									},
								},
							},
							Description: "The Security of NodeConfig.",
						},
						"initialize_script": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The initializeScript of NodeConfig.",
						},
					},
				},
				Description: "The Config of NodePool.",
			},
		},
	}
}

func resourceVolcengineDefaultNodePoolCreate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewDefaultNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Create(nodePoolService, d, ResourceVolcengineDefaultNodePool())
	if err != nil {
		return fmt.Errorf("error on creating DefaultNodePool  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolRead(d, meta)
}

func resourceVolcengineDefaultNodePoolRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewDefaultNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Read(nodePoolService, d, ResourceVolcengineDefaultNodePool())
	if err != nil {
		return fmt.Errorf("error on reading DefaultNodePool %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineDefaultNodePoolUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewDefaultNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Update(nodePoolService, d, ResourceVolcengineDefaultNodePool())
	if err != nil {
		return fmt.Errorf("error on updating DefaultNodePool  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolRead(d, meta)
}

func resourceVolcengineNodePoolDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewDefaultNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Delete(nodePoolService, d, ResourceVolcengineDefaultNodePool())
	if err != nil {
		return fmt.Errorf("error on deleting DefaultNodePool %q, %w", d.Id(), err)
	}
	return err
}
