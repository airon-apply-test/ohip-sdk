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
 * Simple type for Room Condition assignment type. Valid values are Available and NotAvailable Only.
 * @export
 */
export declare const RoomConditionAssigmentType: {
    readonly Available: "Available";
    readonly NotAvailable: "NotAvailable";
};
export type RoomConditionAssigmentType = typeof RoomConditionAssigmentType[keyof typeof RoomConditionAssigmentType];
export declare function RoomConditionAssigmentTypeFromJSON(json: any): RoomConditionAssigmentType;
export declare function RoomConditionAssigmentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomConditionAssigmentType;
export declare function RoomConditionAssigmentTypeToJSON(value?: RoomConditionAssigmentType | null): any;
