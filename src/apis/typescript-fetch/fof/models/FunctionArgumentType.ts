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
 * This type provided details of a function argument and the value it holds.
 * @export
 * @interface FunctionArgumentType
 */
export interface FunctionArgumentType {
    /**
     * Specifies the name of the function argument.
     * @type {string}
     * @memberof FunctionArgumentType
     */
    name?: string;
    /**
     * Specifies the value held by the function argument.
     * @type {string}
     * @memberof FunctionArgumentType
     */
    value?: string;
    /**
     * Specifies the position of the function argument in the argument list.
     * @type {number}
     * @memberof FunctionArgumentType
     */
    position?: number;
    /**
     * Specifies the datatype of the function argument.
     * @type {string}
     * @memberof FunctionArgumentType
     */
    dataType?: string;
    /**
     * Argument of the function mandatory or not.
     * @type {boolean}
     * @memberof FunctionArgumentType
     */
    required?: boolean;
}

/**
 * Check if a given object implements the FunctionArgumentType interface.
 */
export function instanceOfFunctionArgumentType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FunctionArgumentTypeFromJSON(json: any): FunctionArgumentType {
    return FunctionArgumentTypeFromJSONTyped(json, false);
}

export function FunctionArgumentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FunctionArgumentType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'value': !exists(json, 'value') ? undefined : json['value'],
        'position': !exists(json, 'position') ? undefined : json['position'],
        'dataType': !exists(json, 'dataType') ? undefined : json['dataType'],
        'required': !exists(json, 'required') ? undefined : json['required'],
    };
}

export function FunctionArgumentTypeToJSON(value?: FunctionArgumentType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'value': value.value,
        'position': value.position,
        'dataType': value.dataType,
        'required': value.required,
    };
}

