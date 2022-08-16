package main 

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/q48775533q/terraform-provider-pcghost/pcghost" 
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: pcghost.Provider})
}
