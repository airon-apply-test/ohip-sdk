/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Possible values for the Guest Service Status.
 * @export
 */
export const GuestHousekeepingServiceRequestType = {
    DoNotDisturb: 'DoNotDisturb',
    MakeUpRoom: 'MakeUpRoom',
    NoStatusSelected: 'NoStatusSelected'
} as const;
export type GuestHousekeepingServiceRequestType = typeof GuestHousekeepingServiceRequestType[keyof typeof GuestHousekeepingServiceRequestType];


export function GuestHousekeepingServiceRequestTypeFromJSON(json: any): GuestHousekeepingServiceRequestType {
    return GuestHousekeepingServiceRequestTypeFromJSONTyped(json, false);
}

export function GuestHousekeepingServiceRequestTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuestHousekeepingServiceRequestType {
    return json as GuestHousekeepingServiceRequestType;
}

export function GuestHousekeepingServiceRequestTypeToJSON(value?: GuestHousekeepingServiceRequestType | null): any {
    return value as any;
}

