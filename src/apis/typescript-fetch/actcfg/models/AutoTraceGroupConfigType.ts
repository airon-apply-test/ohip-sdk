/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * A collection of supported list of Auto Trace Groups.
 * @export
 */
export const AutoTraceGroupConfigType = {
    Accounts: 'Accounts',
    Contacts: 'Contacts',
    Blocks: 'Blocks',
    Activities: 'Activities'
} as const;
export type AutoTraceGroupConfigType = typeof AutoTraceGroupConfigType[keyof typeof AutoTraceGroupConfigType];


export function AutoTraceGroupConfigTypeFromJSON(json: any): AutoTraceGroupConfigType {
    return AutoTraceGroupConfigTypeFromJSONTyped(json, false);
}

export function AutoTraceGroupConfigTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AutoTraceGroupConfigType {
    return json as AutoTraceGroupConfigType;
}

export function AutoTraceGroupConfigTypeToJSON(value?: AutoTraceGroupConfigType | null): any {
    return value as any;
}

