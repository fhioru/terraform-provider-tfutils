package json_formatter

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFormatJson() *schema.Resource {
	return &schema.Resource{
		Description: "Formats / beautifies a given JSON string.",
		ReadContext: dataSourceFormatJsonRead,
		Schema: map[string]*schema.Schema{
			"json": {
				Type:        schema.TypeString,
				Description: "A JSON string to format / beautify.",
				Required:    true,
			},
			"indent": {
				Type:        schema.TypeString,
				Description: "The character(s) to use for each indent. The default is two spaces.",
				Optional:    true,
				Default:     "  ",
			},
			"result": {
				Type:        schema.TypeString,
				Description: "A formatted JSON string.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFormatJsonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	result := &bytes.Buffer{}

	if err := json.Indent(result, []byte(d.Get("json").(string)), "", d.Get("indent").(string)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("result", result.String()); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
