/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * 
 * @export
 * @interface InterfaceRightsStatusType
 */
export interface InterfaceRightsStatusType {
    /**
     * Internal code for each allowed right of a Hotel Interface record.
     * @type {number}
     * @memberof InterfaceRightsStatusType
     */
    right?: number;
    /**
     * User defined code for an Interface Right.
     * @type {string}
     * @memberof InterfaceRightsStatusType
     */
    statusCode?: string;
    /**
     * User defined description for an Interface Right.
     * @type {string}
     * @memberof InterfaceRightsStatusType
     */
    description?: string;
    /**
     * Category code of this interface right.
     * @type {string}
     * @memberof InterfaceRightsStatusType
     */
    category?: string;
}

/**
 * Check if a given object implements the InterfaceRightsStatusType interface.
 */
export function instanceOfInterfaceRightsStatusType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InterfaceRightsStatusTypeFromJSON(json: any): InterfaceRightsStatusType {
    return InterfaceRightsStatusTypeFromJSONTyped(json, false);
}

export function InterfaceRightsStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceRightsStatusType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'right': !exists(json, 'right') ? undefined : json['right'],
        'statusCode': !exists(json, 'statusCode') ? undefined : json['statusCode'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'category': !exists(json, 'category') ? undefined : json['category'],
    };
}

export function InterfaceRightsStatusTypeToJSON(value?: InterfaceRightsStatusType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'right': value.right,
        'statusCode': value.statusCode,
        'description': value.description,
        'category': value.category,
    };
}

