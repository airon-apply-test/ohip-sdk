/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Simple type for front office room status instructions to be used in requests for fetching housekeeping rooms. Valid values are Vacant and Occupied.
 * @export
 */
export declare const FrontOfficeRoomStatusType: {
    readonly Vacant: "Vacant";
    readonly Occupied: "Occupied";
};
export type FrontOfficeRoomStatusType = typeof FrontOfficeRoomStatusType[keyof typeof FrontOfficeRoomStatusType];
export declare function FrontOfficeRoomStatusTypeFromJSON(json: any): FrontOfficeRoomStatusType;
export declare function FrontOfficeRoomStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FrontOfficeRoomStatusType;
export declare function FrontOfficeRoomStatusTypeToJSON(value?: FrontOfficeRoomStatusType | null): any;
