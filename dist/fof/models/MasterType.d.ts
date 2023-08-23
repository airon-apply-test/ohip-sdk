/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
