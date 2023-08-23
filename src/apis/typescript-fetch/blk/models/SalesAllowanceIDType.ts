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
 * Unique identifier of a single sales allowance.
 * @export
 * @interface SalesAllowanceIDType
 */
export interface SalesAllowanceIDType {
    /**
     * ID of sales allowance.
     * @type {number}
     * @memberof SalesAllowanceIDType
     */
    ceilingId?: number;
}

/**
 * Check if a given object implements the SalesAllowanceIDType interface.
 */
export function instanceOfSalesAllowanceIDType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SalesAllowanceIDTypeFromJSON(json: any): SalesAllowanceIDType {
    return SalesAllowanceIDTypeFromJSONTyped(json, false);
}

export function SalesAllowanceIDTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SalesAllowanceIDType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ceilingId': !exists(json, 'ceilingId') ? undefined : json['ceilingId'],
    };
}

export function SalesAllowanceIDTypeToJSON(value?: SalesAllowanceIDType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ceilingId': value.ceilingId,
    };
}

