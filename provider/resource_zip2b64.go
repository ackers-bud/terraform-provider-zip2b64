package provider

import (
    "github.com/ackers-bud/terraform-provider-zip2b64/client"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
    return &schema.Resource{
        Create: Create,
        Update: Update,
        Read:   ReadUrl,
        Delete: Delete,

        Schema: map[string]*schema.Schema{
            "base64file": {
                Type:     schema.TypeString,
                Required: true,
            },

            "filename": {
                Description: "the filename",
                Type:        schema.TypeString,
                Required:    true,
            },

            "id": {
                Description: "The ID of this resource.",
                Type:        schema.TypeString,
                Computed:    true,
            },

            "filecontents_base64": {
                Description: "The Returned body base64 encoded",
                Type:        schema.TypeString,
                Computed:    true,
            },
        },
    }
}

func Create(d *schema.ResourceData, meta interface{}) error {
    // CreateDiagnostics := diag.Diagnostics{}

    base64file := d.Get("base64file").(string)
    filenameToExtract := d.Get("filename").(string)

    filecontents_base64, _ := client.ZipExtract(base64file, filenameToExtract)

    _ = d.Set("filecontents_base64", filecontents_base64)

    d.SetId(filenameToExtract)
    return nil
}

func Update(d *schema.ResourceData, meta interface{}) error {

    return ReadUrl(d, meta)
}

func ReadUrl(d *schema.ResourceData, meta interface{}) error {

    return nil
}

func Delete(d *schema.ResourceData, meta interface{}) error {

    d.SetId("")

    return nil
}
