/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Rate Plan Asynchronous Service API
 * APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Indicates that the rate code is locked by the central system and cannot be edited at the property level.
 * @export
 */
export const RateCodeLockStatusType = {
    Unlocked: 'Unlocked',
    External: 'External',
    Property: 'Property',
    Central: 'Central'
} as const;
export type RateCodeLockStatusType = typeof RateCodeLockStatusType[keyof typeof RateCodeLockStatusType];


export function RateCodeLockStatusTypeFromJSON(json: any): RateCodeLockStatusType {
    return RateCodeLockStatusTypeFromJSONTyped(json, false);
}

export function RateCodeLockStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RateCodeLockStatusType {
    return json as RateCodeLockStatusType;
}

export function RateCodeLockStatusTypeToJSON(value?: RateCodeLockStatusType | null): any {
    return value as any;
}

