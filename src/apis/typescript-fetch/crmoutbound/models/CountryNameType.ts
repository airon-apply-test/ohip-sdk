/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Code for a country or a nationality.
 * @export
 * @interface CountryNameType
 */
export interface CountryNameType {
    /**
     * Used for Character Strings, length 0 to 200.
     * @type {string}
     * @memberof CountryNameType
     */
    value?: string;
    /**
     * Code for a country or a nationality.
     * @type {string}
     * @memberof CountryNameType
     */
    code?: string;
}

/**
 * Check if a given object implements the CountryNameType interface.
 */
export function instanceOfCountryNameType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CountryNameTypeFromJSON(json: any): CountryNameType {
    return CountryNameTypeFromJSONTyped(json, false);
}

export function CountryNameTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CountryNameType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'value': !exists(json, 'value') ? undefined : json['value'],
        'code': !exists(json, 'code') ? undefined : json['code'],
    };
}

export function CountryNameTypeToJSON(value?: CountryNameType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'value': value.value,
        'code': value.code,
    };
}

