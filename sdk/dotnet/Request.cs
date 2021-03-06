// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Http
{
    [HttpResourceType("http:index:Request")]
    public partial class Request : Pulumi.CustomResource
    {
        [Output("create")]
        public Output<Outputs.Request?> Create { get; private set; } = null!;

        [Output("delete")]
        public Output<Outputs.Request?> Delete { get; private set; } = null!;

        [Output("response")]
        public Output<Outputs.Response> Response { get; private set; } = null!;


        /// <summary>
        /// Create a Request resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Request(string name, RequestArgs? args = null, CustomResourceOptions? options = null)
            : base("http:index:Request", name, args ?? new RequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Request(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("http:index:Request", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Request resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Request Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Request(name, id, options);
        }
    }

    public sealed class RequestArgs : Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<Inputs.RequestArgs>? Create { get; set; }

        [Input("delete")]
        public Input<Inputs.RequestArgs>? Delete { get; set; }

        public RequestArgs()
        {
        }
    }
}
