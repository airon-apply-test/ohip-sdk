/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Instruction to validate changes in reservation and returns validated inventory item.
 * @export
 */
export const ReservationValidationInstructionType = {
    StayHeader: 'StayHeader',
    StayDates: 'StayDates',
    OverlappingStay: 'OverlappingStay',
    Packages: 'Packages',
    RefreshPackages: 'RefreshPackages',
    Preferences: 'Preferences',
    AccessRestrictions: 'AccessRestrictions',
    Ecertificates: 'Ecertificates',
    Promotions: 'Promotions',
    TrackItItems: 'TrackItItems',
    TotalCostOfStay: 'TotalCostOfStay',
    ReservationAwards: 'ReservationAwards',
    WebRegistrationCards: 'WebRegistrationCards',
    ServiceRequests: 'ServiceRequests',
    ReservationActivities: 'ReservationActivities',
    ScheduledActivities: 'ScheduledActivities',
    PrepaidCards: 'PrepaidCards',
    Shares: 'Shares',
    Attachments: 'Attachments',
    Locators: 'Locators',
    ProfileAwards: 'ProfileAwards',
    RefreshRates: 'RefreshRates',
    ReservationGuarantee: 'ReservationGuarantee',
    UpsellsAvailable: 'UpsellsAvailable',
    InventoryItems: 'InventoryItems'
} as const;
export type ReservationValidationInstructionType = typeof ReservationValidationInstructionType[keyof typeof ReservationValidationInstructionType];


export function ReservationValidationInstructionTypeFromJSON(json: any): ReservationValidationInstructionType {
    return ReservationValidationInstructionTypeFromJSONTyped(json, false);
}

export function ReservationValidationInstructionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationValidationInstructionType {
    return json as ReservationValidationInstructionType;
}

export function ReservationValidationInstructionTypeToJSON(value?: ReservationValidationInstructionType | null): any {
    return value as any;
}

