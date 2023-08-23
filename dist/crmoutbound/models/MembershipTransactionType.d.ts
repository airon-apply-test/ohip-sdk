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
import type { MembershipDetailsType } from './MembershipDetailsType';
import type { MembershipTransactionDetailsType } from './MembershipTransactionDetailsType';
/**
 * Detailed transaction information related to a membership.
 * @export
 * @interface MembershipTransactionType
 */
export interface MembershipTransactionType {
    /**
     *
     * @type {MembershipDetailsType}
     * @memberof MembershipTransactionType
     */
    membershipDetails?: MembershipDetailsType;
    /**
     *
     * @type {MembershipTransactionDetailsType}
     * @memberof MembershipTransactionType
     */
    membershipTransactionDetails?: MembershipTransactionDetailsType;
}
/**
 * Check if a given object implements the MembershipTransactionType interface.
 */
export declare function instanceOfMembershipTransactionType(value: object): boolean;
export declare function MembershipTransactionTypeFromJSON(json: any): MembershipTransactionType;
export declare function MembershipTransactionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTransactionType;
export declare function MembershipTransactionTypeToJSON(value?: MembershipTransactionType | null): any;
