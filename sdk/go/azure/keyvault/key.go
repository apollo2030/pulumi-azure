// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Key Vault Key.
type Key struct {
	s *pulumi.ResourceState
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOpt) (*Key, error) {
	if args == nil || args.KeyOpts == nil {
		return nil, errors.New("missing required argument 'KeyOpts'")
	}
	if args == nil || args.KeySize == nil {
		return nil, errors.New("missing required argument 'KeySize'")
	}
	if args == nil || args.KeyType == nil {
		return nil, errors.New("missing required argument 'KeyType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keyOpts"] = nil
		inputs["keySize"] = nil
		inputs["keyType"] = nil
		inputs["keyVaultId"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
		inputs["vaultUri"] = nil
	} else {
		inputs["keyOpts"] = args.KeyOpts
		inputs["keySize"] = args.KeySize
		inputs["keyType"] = args.KeyType
		inputs["keyVaultId"] = args.KeyVaultId
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
		inputs["vaultUri"] = args.VaultUri
	}
	inputs["e"] = nil
	inputs["n"] = nil
	inputs["version"] = nil
	s, err := ctx.RegisterResource("azure:keyvault/key:Key", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Key{s: s}, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyState, opts ...pulumi.ResourceOpt) (*Key, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["e"] = state.E
		inputs["keyOpts"] = state.KeyOpts
		inputs["keySize"] = state.KeySize
		inputs["keyType"] = state.KeyType
		inputs["keyVaultId"] = state.KeyVaultId
		inputs["n"] = state.N
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
		inputs["vaultUri"] = state.VaultUri
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("azure:keyvault/key:Key", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Key{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Key) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Key) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The RSA public exponent of this Key Vault Key.
func (r *Key) E() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["e"])
}

// A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
func (r *Key) KeyOpts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["keyOpts"])
}

// Specifies the Size of the Key to create in bytes. For example, 1024 or 2048. Changing this forces a new resource to be created.
func (r *Key) KeySize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["keySize"])
}

// Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
func (r *Key) KeyType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyType"])
}

// The ID of the Key Vault where the Key should be created.
func (r *Key) KeyVaultId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyVaultId"])
}

// The RSA modulus of this Key Vault Key.
func (r *Key) N() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["n"])
}

// Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
func (r *Key) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A mapping of tags to assign to the resource.
func (r *Key) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

func (r *Key) VaultUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vaultUri"])
}

// The current version of the Key Vault Key.
func (r *Key) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Key resources.
type KeyState struct {
	// The RSA public exponent of this Key Vault Key.
	E interface{}
	// A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
	KeyOpts interface{}
	// Specifies the Size of the Key to create in bytes. For example, 1024 or 2048. Changing this forces a new resource to be created.
	KeySize interface{}
	// Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
	KeyType interface{}
	// The ID of the Key Vault where the Key should be created.
	KeyVaultId interface{}
	// The RSA modulus of this Key Vault Key.
	N interface{}
	// Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	VaultUri interface{}
	// The current version of the Key Vault Key.
	Version interface{}
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
	KeyOpts interface{}
	// Specifies the Size of the Key to create in bytes. For example, 1024 or 2048. Changing this forces a new resource to be created.
	KeySize interface{}
	// Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
	KeyType interface{}
	// The ID of the Key Vault where the Key should be created.
	KeyVaultId interface{}
	// Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	VaultUri interface{}
}
