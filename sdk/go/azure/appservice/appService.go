// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an App Service (within an App Service Plan).
// 
// > **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azurerm_app_service` resource will be overwritten when promoting a Slot using the `azurerm_app_service_active_slot` resource.
type AppService struct {
	s *pulumi.ResourceState
}

// NewAppService registers a new resource with the given unique name, arguments, and options.
func NewAppService(ctx *pulumi.Context,
	name string, args *AppServiceArgs, opts ...pulumi.ResourceOpt) (*AppService, error) {
	if args == nil || args.AppServicePlanId == nil {
		return nil, errors.New("missing required argument 'AppServicePlanId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appServicePlanId"] = nil
		inputs["appSettings"] = nil
		inputs["clientAffinityEnabled"] = nil
		inputs["clientCertEnabled"] = nil
		inputs["connectionStrings"] = nil
		inputs["enabled"] = nil
		inputs["httpsOnly"] = nil
		inputs["identity"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["siteConfig"] = nil
		inputs["tags"] = nil
	} else {
		inputs["appServicePlanId"] = args.AppServicePlanId
		inputs["appSettings"] = args.AppSettings
		inputs["clientAffinityEnabled"] = args.ClientAffinityEnabled
		inputs["clientCertEnabled"] = args.ClientCertEnabled
		inputs["connectionStrings"] = args.ConnectionStrings
		inputs["enabled"] = args.Enabled
		inputs["httpsOnly"] = args.HttpsOnly
		inputs["identity"] = args.Identity
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["siteConfig"] = args.SiteConfig
		inputs["tags"] = args.Tags
	}
	inputs["defaultSiteHostname"] = nil
	inputs["outboundIpAddresses"] = nil
	inputs["possibleOutboundIpAddresses"] = nil
	inputs["siteCredential"] = nil
	inputs["sourceControl"] = nil
	s, err := ctx.RegisterResource("azure:appservice/appService:AppService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppService{s: s}, nil
}

// GetAppService gets an existing AppService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AppServiceState, opts ...pulumi.ResourceOpt) (*AppService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appServicePlanId"] = state.AppServicePlanId
		inputs["appSettings"] = state.AppSettings
		inputs["clientAffinityEnabled"] = state.ClientAffinityEnabled
		inputs["clientCertEnabled"] = state.ClientCertEnabled
		inputs["connectionStrings"] = state.ConnectionStrings
		inputs["defaultSiteHostname"] = state.DefaultSiteHostname
		inputs["enabled"] = state.Enabled
		inputs["httpsOnly"] = state.HttpsOnly
		inputs["identity"] = state.Identity
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["outboundIpAddresses"] = state.OutboundIpAddresses
		inputs["possibleOutboundIpAddresses"] = state.PossibleOutboundIpAddresses
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["siteConfig"] = state.SiteConfig
		inputs["siteCredential"] = state.SiteCredential
		inputs["sourceControl"] = state.SourceControl
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:appservice/appService:AppService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AppService) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AppService) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the App Service Plan within which to create this App Service.
func (r *AppService) AppServicePlanId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appServicePlanId"])
}

// A key-value pair of App Settings.
func (r *AppService) AppSettings() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["appSettings"])
}

// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
func (r *AppService) ClientAffinityEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["clientAffinityEnabled"])
}

// Does the App Service require client certificates for incoming requests? Defaults to `false`.
func (r *AppService) ClientCertEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["clientCertEnabled"])
}

// One or more `connection_string` blocks as defined below.
func (r *AppService) ConnectionStrings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["connectionStrings"])
}

// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
func (r *AppService) DefaultSiteHostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultSiteHostname"])
}

// Is the App Service Enabled?
func (r *AppService) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// Can the App Service only be accessed via HTTPS? Defaults to `false`.
func (r *AppService) HttpsOnly() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["httpsOnly"])
}

// A Managed Service Identity block as defined below.
func (r *AppService) Identity() *pulumi.Output {
	return r.s.State["identity"]
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *AppService) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the App Service. Changing this forces a new resource to be created.
func (r *AppService) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
func (r *AppService) OutboundIpAddresses() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["outboundIpAddresses"])
}

// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
func (r *AppService) PossibleOutboundIpAddresses() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["possibleOutboundIpAddresses"])
}

// The name of the resource group in which to create the App Service.
func (r *AppService) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `site_config` block as defined below.
func (r *AppService) SiteConfig() *pulumi.Output {
	return r.s.State["siteConfig"]
}

// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
func (r *AppService) SiteCredential() *pulumi.Output {
	return r.s.State["siteCredential"]
}

// A `source_control` block as defined below, which contains the Source Control information when `scm_type` is set to `LocalGit`.
func (r *AppService) SourceControl() *pulumi.Output {
	return r.s.State["sourceControl"]
}

// A mapping of tags to assign to the resource.
func (r *AppService) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering AppService resources.
type AppServiceState struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId interface{}
	// A key-value pair of App Settings.
	AppSettings interface{}
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled interface{}
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled interface{}
	// One or more `connection_string` blocks as defined below.
	ConnectionStrings interface{}
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname interface{}
	// Is the App Service Enabled?
	Enabled interface{}
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly interface{}
	// A Managed Service Identity block as defined below.
	Identity interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name interface{}
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses interface{}
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
	PossibleOutboundIpAddresses interface{}
	// The name of the resource group in which to create the App Service.
	ResourceGroupName interface{}
	// A `site_config` block as defined below.
	SiteConfig interface{}
	// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredential interface{}
	// A `source_control` block as defined below, which contains the Source Control information when `scm_type` is set to `LocalGit`.
	SourceControl interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a AppService resource.
type AppServiceArgs struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId interface{}
	// A key-value pair of App Settings.
	AppSettings interface{}
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled interface{}
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled interface{}
	// One or more `connection_string` blocks as defined below.
	ConnectionStrings interface{}
	// Is the App Service Enabled?
	Enabled interface{}
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly interface{}
	// A Managed Service Identity block as defined below.
	Identity interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the App Service.
	ResourceGroupName interface{}
	// A `site_config` block as defined below.
	SiteConfig interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
