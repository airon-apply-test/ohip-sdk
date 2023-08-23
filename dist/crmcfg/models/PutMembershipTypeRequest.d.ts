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
import type { MembershipTypeChangeInstructionType } from './MembershipTypeChangeInstructionType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PutMembershipTypeRequest
 */
export interface PutMembershipTypeRequest {
    /**
     * A collection of MembershipTypes with information that needs to be changed.
     * @type {Array<MembershipTypeChangeInstructionType>}
     * @memberof PutMembershipTypeRequest
     */
    membershipTypeChangeInstructions?: Array<MembershipTypeChangeInstructionType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PutMembershipTypeRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutMembershipTypeRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PutMembershipTypeRequest interface.
 */
export declare function instanceOfPutMembershipTypeRequest(value: object): boolean;
export declare function PutMembershipTypeRequestFromJSON(json: any): PutMembershipTypeRequest;
export declare function PutMembershipTypeRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutMembershipTypeRequest;
export declare function PutMembershipTypeRequestToJSON(value?: PutMembershipTypeRequest | null): any;
