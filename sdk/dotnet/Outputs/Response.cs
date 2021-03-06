// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Http.Outputs
{

    [OutputType]
    public sealed class Response
    {
        public readonly string Body;
        public readonly Outputs.Header Header;
        public readonly string Status;
        public readonly int StatusCode;

        [OutputConstructor]
        private Response(
            string body,

            Outputs.Header header,

            string status,

            int statusCode)
        {
            Body = body;
            Header = header;
            Status = status;
            StatusCode = statusCode;
        }
    }
}
