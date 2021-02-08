// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Auth0.Outputs
{

    [OutputType]
    public sealed class ConnectionOptionsMfa
    {
        /// <summary>
        /// Indicates whether multifactor authentication is enabled for this connection.
        /// </summary>
        public readonly bool? Active;
        /// <summary>
        /// Indicates whether multifactor authentication enrollment settings will be returned.
        /// </summary>
        public readonly bool? ReturnEnrollSettings;

        [OutputConstructor]
        private ConnectionOptionsMfa(
            bool? active,

            bool? returnEnrollSettings)
        {
            Active = active;
            ReturnEnrollSettings = returnEnrollSettings;
        }
    }
}