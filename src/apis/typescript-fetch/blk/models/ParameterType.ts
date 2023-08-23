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
 * Name value pair type that will hold generic parameter information. Only use this type when the parameters being passed are too dynamic to be defined.
 * @export
 * @interface ParameterType
 */
export interface ParameterType {
    /**
     * Name of the parameter.
     * @type {string}
     * @memberof ParameterType
     */
    parameterName?: string;
    /**
     * Value of the parameter.
     * @type {string}
     * @memberof ParameterType
     */
    parameterValue?: string;
}

/**
 * Check if a given object implements the ParameterType interface.
 */
export function instanceOfParameterType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ParameterTypeFromJSON(json: any): ParameterType {
    return ParameterTypeFromJSONTyped(json, false);
}

export function ParameterTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ParameterType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'parameterName': !exists(json, 'parameterName') ? undefined : json['parameterName'],
        'parameterValue': !exists(json, 'parameterValue') ? undefined : json['parameterValue'],
    };
}

export function ParameterTypeToJSON(value?: ParameterType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'parameterName': value.parameterName,
        'parameterValue': value.parameterValue,
    };
}

