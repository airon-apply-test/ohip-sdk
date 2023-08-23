/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * This reservation is in checked in status and the business date is past departure date. This could occur when ORS and PMS are in same environment.
 * @export
 */
export declare const PMSResStatusType: {
    readonly Reserved: "Reserved";
    readonly Requested: "Requested";
    readonly NoShow: "NoShow";
    readonly Cancelled: "Cancelled";
    readonly InHouse: "InHouse";
    readonly CheckedOut: "CheckedOut";
    readonly Waitlisted: "Waitlisted";
    readonly DueIn: "DueIn";
    readonly DueOut: "DueOut";
    readonly Walkin: "Walkin";
    readonly PendingCheckout: "PendingCheckout";
};
export type PMSResStatusType = typeof PMSResStatusType[keyof typeof PMSResStatusType];
export declare function PMSResStatusTypeFromJSON(json: any): PMSResStatusType;
export declare function PMSResStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PMSResStatusType;
export declare function PMSResStatusTypeToJSON(value?: PMSResStatusType | null): any;
