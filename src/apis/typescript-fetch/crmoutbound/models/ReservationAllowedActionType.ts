/* tslint:disable */
/* eslint-disable */
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
 * Flag indicating whether the Enrollment is in progress or not for the Profile associated with this Reservation.
 * @export
 */
export const ReservationAllowedActionType = {
    Cancel: 'Cancel',
    Move: 'Move',
    PreCharge: 'PreCharge',
    PostCharge: 'PostCharge',
    FacilitySchedule: 'FacilitySchedule',
    Upsell: 'Upsell',
    PreCheckIn: 'PreCheckIn',
    PostToNoShowCancel: 'PostToNoShowCancel',
    NoShow: 'NoShow',
    NameChange: 'NameChange',
    Discount: 'Discount',
    EnrollToPrimaryMembership: 'EnrollToPrimaryMembership',
    EnrollInProgress: 'EnrollInProgress'
} as const;
export type ReservationAllowedActionType = typeof ReservationAllowedActionType[keyof typeof ReservationAllowedActionType];


export function ReservationAllowedActionTypeFromJSON(json: any): ReservationAllowedActionType {
    return ReservationAllowedActionTypeFromJSONTyped(json, false);
}

export function ReservationAllowedActionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationAllowedActionType {
    return json as ReservationAllowedActionType;
}

export function ReservationAllowedActionTypeToJSON(value?: ReservationAllowedActionType | null): any {
    return value as any;
}

