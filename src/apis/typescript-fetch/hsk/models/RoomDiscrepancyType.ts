/* tslint:disable */
/* eslint-disable */
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
 * Simple type for discrepant room status instructions to be used in requests for fetching housekeeping and front office room discrepancy. Valid values are Sleep, Skip, Person Discrepancy, Due Out Only.
 * @export
 */
export const RoomDiscrepancyType = {
    Sleep: 'Sleep',
    Skip: 'Skip',
    PersonDiscrepancy: 'PersonDiscrepancy'
} as const;
export type RoomDiscrepancyType = typeof RoomDiscrepancyType[keyof typeof RoomDiscrepancyType];


export function RoomDiscrepancyTypeFromJSON(json: any): RoomDiscrepancyType {
    return RoomDiscrepancyTypeFromJSONTyped(json, false);
}

export function RoomDiscrepancyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomDiscrepancyType {
    return json as RoomDiscrepancyType;
}

export function RoomDiscrepancyTypeToJSON(value?: RoomDiscrepancyType | null): any {
    return value as any;
}

