/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Message Waiting is set to available and freely switchable.
 * @export
 */
export const MessageWaitingType = {
    NotAvailable: 'NotAvailable',
    AlwaysOn: 'AlwaysOn',
    AlwaysOff: 'AlwaysOff',
    FreelySwitchable: 'FreelySwitchable'
} as const;
export type MessageWaitingType = typeof MessageWaitingType[keyof typeof MessageWaitingType];


export function MessageWaitingTypeFromJSON(json: any): MessageWaitingType {
    return MessageWaitingTypeFromJSONTyped(json, false);
}

export function MessageWaitingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageWaitingType {
    return json as MessageWaitingType;
}

export function MessageWaitingTypeToJSON(value?: MessageWaitingType | null): any {
    return value as any;
}

