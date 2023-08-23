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
 * Indicates any Updates/Changes on the reservation as well as Reservation Cancellation are not allowed.
 * @export
 */
export const ResAccessRestrictionType = {
    Change: 'Change',
    Cancel: 'Cancel',
    ChangeOrCancel: 'ChangeOrCancel'
} as const;
export type ResAccessRestrictionType = typeof ResAccessRestrictionType[keyof typeof ResAccessRestrictionType];


export function ResAccessRestrictionTypeFromJSON(json: any): ResAccessRestrictionType {
    return ResAccessRestrictionTypeFromJSONTyped(json, false);
}

export function ResAccessRestrictionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResAccessRestrictionType {
    return json as ResAccessRestrictionType;
}

export function ResAccessRestrictionTypeToJSON(value?: ResAccessRestrictionType | null): any {
    return value as any;
}

