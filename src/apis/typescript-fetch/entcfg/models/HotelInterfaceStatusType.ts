/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Status of the Hotel Interface either STOPPED or RUNNING.
 * @export
 */
export const HotelInterfaceStatusType = {
    Stopped: 'Stopped',
    Running: 'Running',
    Waiting: 'Waiting',
    StopInitiated: 'StopInitiated',
    StartInitiated: 'StartInitiated',
    RebootInitiated: 'RebootInitiated',
    Other: 'Other'
} as const;
export type HotelInterfaceStatusType = typeof HotelInterfaceStatusType[keyof typeof HotelInterfaceStatusType];


export function HotelInterfaceStatusTypeFromJSON(json: any): HotelInterfaceStatusType {
    return HotelInterfaceStatusTypeFromJSONTyped(json, false);
}

export function HotelInterfaceStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInterfaceStatusType {
    return json as HotelInterfaceStatusType;
}

export function HotelInterfaceStatusTypeToJSON(value?: HotelInterfaceStatusType | null): any {
    return value as any;
}

