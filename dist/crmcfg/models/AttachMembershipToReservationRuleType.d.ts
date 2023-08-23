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
/**
 * Indicates when the membership information should get attach to the reservation.
 * @export
 */
export declare const AttachMembershipToReservationRuleType: {
    readonly NewReservation: "NewReservation";
    readonly UpdateReservation: "UpdateReservation";
    readonly Checkin: "Checkin";
    readonly Checkout: "Checkout";
};
export type AttachMembershipToReservationRuleType = typeof AttachMembershipToReservationRuleType[keyof typeof AttachMembershipToReservationRuleType];
export declare function AttachMembershipToReservationRuleTypeFromJSON(json: any): AttachMembershipToReservationRuleType;
export declare function AttachMembershipToReservationRuleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AttachMembershipToReservationRuleType;
export declare function AttachMembershipToReservationRuleTypeToJSON(value?: AttachMembershipToReservationRuleType | null): any;
