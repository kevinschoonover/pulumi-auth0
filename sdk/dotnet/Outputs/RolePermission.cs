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
    public sealed class RolePermission
    {
        /// <summary>
        /// String. Name of the permission (scope).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// String. Unique identifier for the resource server.
        /// </summary>
        public readonly string ResourceServerIdentifier;

        [OutputConstructor]
        private RolePermission(
            string name,

            string resourceServerIdentifier)
        {
            Name = name;
            ResourceServerIdentifier = resourceServerIdentifier;
        }
    }
}