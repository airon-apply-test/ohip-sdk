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
 * The supplier's ranking of the customer.
 * @export
 * @interface VIPType
 */
export interface VIPType {
    /**
     * VIP Code.
     * @type {string}
     * @memberof VIPType
     */
    vipCode?: string;
    /**
     * VIP Description.
     * @type {string}
     * @memberof VIPType
     */
    vipDescription?: string;
}

/**
 * Check if a given object implements the VIPType interface.
 */
export function instanceOfVIPType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function VIPTypeFromJSON(json: any): VIPType {
    return VIPTypeFromJSONTyped(json, false);
}

export function VIPTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): VIPType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'vipCode': !exists(json, 'vipCode') ? undefined : json['vipCode'],
        'vipDescription': !exists(json, 'vipDescription') ? undefined : json['vipDescription'],
    };
}

export function VIPTypeToJSON(value?: VIPType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'vipCode': value.vipCode,
        'vipDescription': value.vipDescription,
    };
}

