/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Search match indicating attribute and the matching value.
 * @export
 * @interface SearchMatchType
 */
export interface SearchMatchType {
    /**
     * Search match attribute.
     * @type {string}
     * @memberof SearchMatchType
     */
    attribute?: string;
    /**
     * Search match value.
     * @type {string}
     * @memberof SearchMatchType
     */
    value?: string;
}

/**
 * Check if a given object implements the SearchMatchType interface.
 */
export function instanceOfSearchMatchType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SearchMatchTypeFromJSON(json: any): SearchMatchType {
    return SearchMatchTypeFromJSONTyped(json, false);
}

export function SearchMatchTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchMatchType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'attribute': !exists(json, 'attribute') ? undefined : json['attribute'],
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function SearchMatchTypeToJSON(value?: SearchMatchType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'attribute': value.attribute,
        'value': value.value,
    };
}

