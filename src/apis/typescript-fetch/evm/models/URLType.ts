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

import { exists, mapValues } from '../runtime';
/**
 * Web site address, in IETF(The Internet Engineering Task Force) specified format.
 * @export
 * @interface URLType
 */
export interface URLType {
    /**
     * Property Value
     * @type {string}
     * @memberof URLType
     */
    value?: string;
    /**
     * Defines the purpose of the URL address, such as personal, business, public, etc.
     * @type {string}
     * @memberof URLType
     */
    type?: string;
    /**
     * Describes the Type code
     * @type {string}
     * @memberof URLType
     */
    typeDescription?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof URLType
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof URLType
     */
    orderSequence?: number;
}

/**
 * Check if a given object implements the URLType interface.
 */
export function instanceOfURLType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function URLTypeFromJSON(json: any): URLType {
    return URLTypeFromJSONTyped(json, false);
}

export function URLTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): URLType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'value': !exists(json, 'value') ? undefined : json['value'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'typeDescription': !exists(json, 'typeDescription') ? undefined : json['typeDescription'],
        'primaryInd': !exists(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}

export function URLTypeToJSON(value?: URLType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'value': value.value,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
    };
}

