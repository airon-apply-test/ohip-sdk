/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Type of the entity that is requesting for the item hold.
 * @export
 */
export const HeldByType = {
    Event: 'Event',
    Reservation: 'Reservation',
    Allotment: 'Allotment'
} as const;
export type HeldByType = typeof HeldByType[keyof typeof HeldByType];


export function HeldByTypeFromJSON(json: any): HeldByType {
    return HeldByTypeFromJSONTyped(json, false);
}

export function HeldByTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HeldByType {
    return json as HeldByType;
}

export function HeldByTypeToJSON(value?: HeldByType | null): any {
    return value as any;
}

