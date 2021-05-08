package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/lucasepe/terraform-provider-codename/internal/codename"
)

func resourceCodename() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCodename,
		ReadContext:   schema.NoopContext,
		DeleteContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
			d.SetId("")
			return nil
		},

		Schema: map[string]*schema.Schema{
			"token_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
			},

			"prefix": {
				Description: "A string to prefix the coedname with.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},

			"snakefy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"id": {
				Description: "The generated codename",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func createCodename(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	rng := meta.(*rand.Rand)
	if rng == nil {
		err := fmt.Errorf("invalid source of random numbers: [%v]", rng)
		return append(diags, diag.Errorf("Error RNG: %s", err)...)
	}

	prefix := d.Get("prefix").(string)
	tokenLength := d.Get("token_length").(int)
	snakefy := d.Get("snakefy").(bool)

	codename := codename.Generate(rng, tokenLength, snakefy)
	// log.Printf("[DEBUG] ---> %s <----\n", codename)
	// ^^^^ just to show how to debug using TF_LOG=TRACE

	if prefix != "" {
		codename = fmt.Sprintf("%s%s", prefix, codename)
	}

	// sets the ID of the resource otherwise the resource is destroyed.
	d.SetId(codename)

	return diags
}
