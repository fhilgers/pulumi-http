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
    public sealed class Request
    {
        public readonly string? Body;
        public readonly ImmutableArray<string> Certificates;
        public readonly double? ExpectedStatusCode;
        public readonly Outputs.Header? Header;
        public readonly bool? InsecureSkipVerify;
        public readonly double? MaxRetries;
        public readonly string Method;
        public readonly double? RetryWaitMax;
        public readonly double? RetryWaitMin;
        public readonly ImmutableArray<string> RootCAs;
        public readonly string? ServerName;
        public readonly string Url;

        [OutputConstructor]
        private Request(
            string? body,

            ImmutableArray<string> certificates,

            double? expectedStatusCode,

            Outputs.Header? header,

            bool? insecureSkipVerify,

            double? maxRetries,

            string method,

            double? retryWaitMax,

            double? retryWaitMin,

            ImmutableArray<string> rootCAs,

            string? serverName,

            string url)
        {
            Body = body;
            Certificates = certificates;
            ExpectedStatusCode = expectedStatusCode;
            Header = header;
            InsecureSkipVerify = insecureSkipVerify;
            MaxRetries = maxRetries;
            Method = method;
            RetryWaitMax = retryWaitMax;
            RetryWaitMin = retryWaitMin;
            RootCAs = rootCAs;
            ServerName = serverName;
            Url = url;
        }
    }
}
