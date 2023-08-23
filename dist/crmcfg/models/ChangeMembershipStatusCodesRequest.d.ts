/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { MembershipStatusCodeType } from './MembershipStatusCodeType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface ChangeMembershipStatusCodesRequest
 */
export interface ChangeMembershipStatusCodesRequest {
    /**
     * List of Membership Status Codes.
     * @type {Array<MembershipStatusCodeType>}
     * @memberof ChangeMembershipStatusCodesRequest
     */
    membershipStatusCodes?: Array<MembershipStatusCodeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ChangeMembershipStatusCodesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChangeMembershipStatusCodesRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ChangeMembershipStatusCodesRequest interface.
 */
export declare function instanceOfChangeMembershipStatusCodesRequest(value: object): boolean;
export declare function ChangeMembershipStatusCodesRequestFromJSON(json: any): ChangeMembershipStatusCodesRequest;
export declare function ChangeMembershipStatusCodesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeMembershipStatusCodesRequest;
export declare function ChangeMembershipStatusCodesRequestToJSON(value?: ChangeMembershipStatusCodesRequest | null): any;
