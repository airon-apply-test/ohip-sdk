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
import type { MembershipEnrollmentCodeType } from './MembershipEnrollmentCodeType';
import type { WarningType } from './WarningType';
/**
 * Request object for changing Membership Enrollment Codes.
 * @export
 * @interface MembershipEnrollmentCodesToBeChanged
 */
export interface MembershipEnrollmentCodesToBeChanged {
    /**
     * List of Membership Enrollment Codes.
     * @type {Array<MembershipEnrollmentCodeType>}
     * @memberof MembershipEnrollmentCodesToBeChanged
     */
    membershipEnrollmentCodes?: Array<MembershipEnrollmentCodeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof MembershipEnrollmentCodesToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof MembershipEnrollmentCodesToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the MembershipEnrollmentCodesToBeChanged interface.
 */
export declare function instanceOfMembershipEnrollmentCodesToBeChanged(value: object): boolean;
export declare function MembershipEnrollmentCodesToBeChangedFromJSON(json: any): MembershipEnrollmentCodesToBeChanged;
export declare function MembershipEnrollmentCodesToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipEnrollmentCodesToBeChanged;
export declare function MembershipEnrollmentCodesToBeChangedToJSON(value?: MembershipEnrollmentCodesToBeChanged | null): any;
