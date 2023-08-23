/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { MembershipClaimDetailsType } from './MembershipClaimDetailsType';
/**
 * Request object for creating membership claim.
 * @export
 * @interface MembershipClaim
 */
export interface MembershipClaim {
    /**
     *
     * @type {MembershipClaimDetailsType}
     * @memberof MembershipClaim
     */
    membershipClaimDetails?: MembershipClaimDetailsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof MembershipClaim
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the MembershipClaim interface.
 */
export declare function instanceOfMembershipClaim(value: object): boolean;
export declare function MembershipClaimFromJSON(json: any): MembershipClaim;
export declare function MembershipClaimFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipClaim;
export declare function MembershipClaimToJSON(value?: MembershipClaim | null): any;
