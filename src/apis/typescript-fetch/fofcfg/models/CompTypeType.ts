/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Information representation of Comp Type.
 * @export
 * @interface CompTypeType
 */
export interface CompTypeType {
    /**
     * Code is used to identify the CompType.
     * @type {string}
     * @memberof CompTypeType
     */
    compTypeCode?: string;
    /**
     * Text is used to describe the CompType.
     * @type {string}
     * @memberof CompTypeType
     */
    compTypeDescription?: string;
    /**
     * Comp Type record sequence number.
     * @type {number}
     * @memberof CompTypeType
     */
    displayOrder?: number;
}

/**
 * Check if a given object implements the CompTypeType interface.
 */
export function instanceOfCompTypeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CompTypeTypeFromJSON(json: any): CompTypeType {
    return CompTypeTypeFromJSONTyped(json, false);
}

export function CompTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompTypeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'compTypeCode': !exists(json, 'compTypeCode') ? undefined : json['compTypeCode'],
        'compTypeDescription': !exists(json, 'compTypeDescription') ? undefined : json['compTypeDescription'],
        'displayOrder': !exists(json, 'displayOrder') ? undefined : json['displayOrder'],
    };
}

export function CompTypeTypeToJSON(value?: CompTypeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'compTypeCode': value.compTypeCode,
        'compTypeDescription': value.compTypeDescription,
        'displayOrder': value.displayOrder,
    };
}

