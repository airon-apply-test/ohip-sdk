/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Lamp indicator Type.
 * @export
 * @interface IndicatorType
 */
export interface IndicatorType {
    /**
     * Name of the indicator.
     * @type {string}
     * @memberof IndicatorType
     */
    indicatorName?: string;
    /**
     * Indicates number of occurrences of the indicator.
     * @type {number}
     * @memberof IndicatorType
     */
    count?: number;
}

/**
 * Check if a given object implements the IndicatorType interface.
 */
export function instanceOfIndicatorType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function IndicatorTypeFromJSON(json: any): IndicatorType {
    return IndicatorTypeFromJSONTyped(json, false);
}

export function IndicatorTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IndicatorType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'indicatorName': !exists(json, 'indicatorName') ? undefined : json['indicatorName'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function IndicatorTypeToJSON(value?: IndicatorType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'indicatorName': value.indicatorName,
        'count': value.count,
    };
}

