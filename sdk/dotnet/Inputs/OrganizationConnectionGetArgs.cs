// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Auth0.Inputs
{

    public sealed class OrganizationConnectionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, all users that log in
        /// with this connection will be automatically granted membership in the
        /// organization. When false, users must be granted membership in the organization
        /// before logging in with this connection.
        /// </summary>
        [Input("assignMembershipOnLogin")]
        public Input<bool>? AssignMembershipOnLogin { get; set; }

        /// <summary>
        /// The connection ID of the connection to add to the
        /// organization
        /// </summary>
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        public OrganizationConnectionGetArgs()
        {
        }
    }
}
