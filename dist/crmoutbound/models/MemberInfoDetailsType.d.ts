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
/**
 *
 * @export
 * @interface MemberInfoDetailsType
 */
export interface MemberInfoDetailsType {
    /**
     * Total nights of the guest.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    guestTotalNights?: number;
    /**
     * STotal stays of the guest.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    guestTotalStay?: number;
    /**
     * Tier Base Nights Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBaseNights?: number;
    /**
     * Tier Base Revenue Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBaseRevenue?: number;
    /**
     * Tier Base Stay Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBaseStay?: number;
    /**
     * Tier Bonus Nights Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBonusNights?: number;
    /**
     * Tier Bonus Revenue Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBonusRevenue?: number;
    /**
     * Tier Bonus Stay Points.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    tierBonusStay?: number;
    /**
     * Total Nights of the membership.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    membershipTotalNights?: number;
    /**
     * Total Stay of the membership.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    membershipTotalStay?: number;
    /**
     * Total base points of the award.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    awardBasePoints?: number;
    /**
     * Total bonus points of the award.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    awardBonusPoints?: number;
    /**
     * Total miscellaneous points of the award.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    awardMiscPoints?: number;
    /**
     * Total points of the award.
     * @type {number}
     * @memberof MemberInfoDetailsType
     */
    totalAwardPoints?: number;
}
/**
 * Check if a given object implements the MemberInfoDetailsType interface.
 */
export declare function instanceOfMemberInfoDetailsType(value: object): boolean;
export declare function MemberInfoDetailsTypeFromJSON(json: any): MemberInfoDetailsType;
export declare function MemberInfoDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MemberInfoDetailsType;
export declare function MemberInfoDetailsTypeToJSON(value?: MemberInfoDetailsType | null): any;
