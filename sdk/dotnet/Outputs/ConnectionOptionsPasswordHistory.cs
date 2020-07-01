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
    public sealed class ConnectionOptionsPasswordHistory
    {
        /// <summary>
        /// Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// Integer, (Maximum=24). Indicates the number of passwords to keep in history.
        /// </summary>
        public readonly int? Size;

        [OutputConstructor]
        private ConnectionOptionsPasswordHistory(
            bool? enable,

            int? size)
        {
            Enable = enable;
            Size = size;
        }
    }
}
