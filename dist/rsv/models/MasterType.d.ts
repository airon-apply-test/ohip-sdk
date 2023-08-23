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
 *
 * @export
 */
export declare const MasterType: {
    readonly Country: "Country";
    readonly State: "State";
    readonly AddressType: "AddressType";
    readonly PhoneType: "PhoneType";
    readonly RateCategory: "RateCategory";
    readonly CalculationRule: "CalculationRule";
    readonly PostingRythym: "PostingRythym";
    readonly BillingInstruction: "BillingInstruction";
    readonly TransactionCode: "TransactionCode";
    readonly DisplaySet: "DisplaySet";
    readonly MailingActions: "MailingActions";
    readonly DistanceType: "DistanceType";
    readonly District: "District";
    readonly Territory: "Territory";
    readonly FiscalRegion: "FiscalRegion";
    readonly InventoryItem: "InventoryItem";
    readonly Package: "Package";
    readonly RoomFeaturePreference: "RoomFeaturePreference";
    readonly SpecialPreference: "SpecialPreference";
    readonly Promotion: "Promotion";
    readonly Department: "Department";
    readonly ReservationPreference: "ReservationPreference";
    readonly FacilityTask: "FacilityTask";
    readonly RoomType: "RoomType";
    readonly RateCode: "RateCode";
    readonly OutOfOrderReason: "OutOfOrderReason";
    readonly Block: "Block";
};
export type MasterType = typeof MasterType[keyof typeof MasterType];
export declare function MasterTypeFromJSON(json: any): MasterType;
export declare function MasterTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MasterType;
export declare function MasterTypeToJSON(value?: MasterType | null): any;
