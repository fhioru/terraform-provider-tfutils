package tfutils

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJsonFormat() *schema.Resource {
	return &schema.Resource{
		Description: "Formats / beautifies a given JSON string.",
		ReadContext: dataSourceJsonFormatRead,
		Schema: map[string]*schema.Schema{
			"json": {
				Type:        schema.TypeString,
				Description: "A JSON string to format / beautify.",
				Required:    true,
			},
			"indent": {
				Type:        schema.TypeInt,
				Description: "The number of spaces to use for each indent. The default is two spaces.",
				Optional:    true,
				Default:     2,
			},
			"result": {
				Type:        schema.TypeString,
				Description: "A formatted JSON string.",
				Optional:    true,
			},
		},
	}
}

func dataSourceJsonFormatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	result := &bytes.Buffer{}

	spaces := strings.Repeat(" ", d.Get("indent").(int))

	if err := json.Indent(result, []byte(d.Get("json").(string)), "", spaces); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("result", result.String()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
