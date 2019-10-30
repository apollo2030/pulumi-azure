// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Logicapps
{
    /// <summary>
    /// Manages an HTTP Action within a Logic App Workflow
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/logic_app_action_http.html.markdown.
    /// </summary>
    public partial class ActionHttp : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        /// <summary>
        /// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        [Output("headers")]
        public Output<ImmutableDictionary<string, string>?> Headers { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("logicAppId")]
        public Output<string> LogicAppId { get; private set; } = null!;

        /// <summary>
        /// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        /// </summary>
        [Output("method")]
        public Output<string> Method { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the URI which will be called when this HTTP Action is triggered.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a ActionHttp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionHttp(string name, ActionHttpArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/actionHttp:ActionHttp", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ActionHttp(string name, Input<string> id, ActionHttpState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/actionHttp:ActionHttp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActionHttp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionHttp Get(string name, Input<string> id, ActionHttpState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionHttp(name, id, state, options);
        }
    }

    public sealed class ActionHttpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId", required: true)]
        public Input<string> LogicAppId { get; set; } = null!;

        /// <summary>
        /// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the URI which will be called when this HTTP Action is triggered.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public ActionHttpArgs()
        {
        }
    }

    public sealed class ActionHttpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId")]
        public Input<string>? LogicAppId { get; set; }

        /// <summary>
        /// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the URI which will be called when this HTTP Action is triggered.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ActionHttpState()
        {
        }
    }
}