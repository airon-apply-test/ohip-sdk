/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AddtionalCodeInfoTypeInner } from './AddtionalCodeInfoTypeInner';
import {
    AddtionalCodeInfoTypeInnerFromJSON,
    AddtionalCodeInfoTypeInnerFromJSONTyped,
    AddtionalCodeInfoTypeInnerToJSON,
} from './AddtionalCodeInfoTypeInner';

/**
 * 
 * @export
 * @interface CodeInfoType
 */
export interface CodeInfoType {
    /**
     * 
     * @type {string}
     * @memberof CodeInfoType
     */
    description?: string;
    /**
     * Holds name of additional code information
     * @type {Array<AddtionalCodeInfoTypeInner>}
     * @memberof CodeInfoType
     */
    addtionalCodeInfo?: Array<AddtionalCodeInfoTypeInner>;
    /**
     * 
     * @type {string}
     * @memberof CodeInfoType
     */
    hotelId?: string;
    /**
     * 
     * @type {string}
     * @memberof CodeInfoType
     */
    code?: string;
}

/**
 * Check if a given object implements the CodeInfoType interface.
 */
export function instanceOfCodeInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CodeInfoTypeFromJSON(json: any): CodeInfoType {
    return CodeInfoTypeFromJSONTyped(json, false);
}

export function CodeInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CodeInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': !exists(json, 'description') ? undefined : json['description'],
        'addtionalCodeInfo': !exists(json, 'addtionalCodeInfo') ? undefined : ((json['addtionalCodeInfo'] as Array<any>).map(AddtionalCodeInfoTypeInnerFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !exists(json, 'code') ? undefined : json['code'],
    };
}

export function CodeInfoTypeToJSON(value?: CodeInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
        'addtionalCodeInfo': value.addtionalCodeInfo === undefined ? undefined : ((value.addtionalCodeInfo as Array<any>).map(AddtionalCodeInfoTypeInnerToJSON)),
        'hotelId': value.hotelId,
        'code': value.code,
    };
}

