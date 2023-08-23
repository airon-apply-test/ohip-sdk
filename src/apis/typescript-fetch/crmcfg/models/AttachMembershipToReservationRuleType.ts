/* tslint:disable */
/* eslint-disable */
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
export const AttachMembershipToReservationRuleType = {
    NewReservation: 'NewReservation',
    UpdateReservation: 'UpdateReservation',
    Checkin: 'Checkin',
    Checkout: 'Checkout'
} as const;
export type AttachMembershipToReservationRuleType = typeof AttachMembershipToReservationRuleType[keyof typeof AttachMembershipToReservationRuleType];


export function AttachMembershipToReservationRuleTypeFromJSON(json: any): AttachMembershipToReservationRuleType {
    return AttachMembershipToReservationRuleTypeFromJSONTyped(json, false);
}

export function AttachMembershipToReservationRuleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AttachMembershipToReservationRuleType {
    return json as AttachMembershipToReservationRuleType;
}

export function AttachMembershipToReservationRuleTypeToJSON(value?: AttachMembershipToReservationRuleType | null): any {
    return value as any;
}

