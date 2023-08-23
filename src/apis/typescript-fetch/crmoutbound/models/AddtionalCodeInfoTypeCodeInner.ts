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
import type { MasterInfoCodeDetailType } from './MasterInfoCodeDetailType';
import {
    MasterInfoCodeDetailTypeFromJSON,
    MasterInfoCodeDetailTypeFromJSONTyped,
    MasterInfoCodeDetailTypeToJSON,
} from './MasterInfoCodeDetailType';

/**
 * 
 * @export
 * @interface AddtionalCodeInfoTypeCodeInner
 */
export interface AddtionalCodeInfoTypeCodeInner {
    /**
     * 
     * @type {MasterInfoCodeDetailType}
     * @memberof AddtionalCodeInfoTypeCodeInner
     */
    name?: MasterInfoCodeDetailType;
    /**
     * Holds value of additional code information
     * @type {string}
     * @memberof AddtionalCodeInfoTypeCodeInner
     */
    value?: string;
}

/**
 * Check if a given object implements the AddtionalCodeInfoTypeCodeInner interface.
 */
export function instanceOfAddtionalCodeInfoTypeCodeInner(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AddtionalCodeInfoTypeCodeInnerFromJSON(json: any): AddtionalCodeInfoTypeCodeInner {
    return AddtionalCodeInfoTypeCodeInnerFromJSONTyped(json, false);
}

export function AddtionalCodeInfoTypeCodeInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddtionalCodeInfoTypeCodeInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : MasterInfoCodeDetailTypeFromJSON(json['name']),
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function AddtionalCodeInfoTypeCodeInnerToJSON(value?: AddtionalCodeInfoTypeCodeInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': MasterInfoCodeDetailTypeToJSON(value.name),
        'value': value.value,
    };
}

