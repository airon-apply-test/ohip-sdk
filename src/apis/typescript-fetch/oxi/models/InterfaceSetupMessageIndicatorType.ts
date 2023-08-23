/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Type represents indicators for deleting/keeping data in the OPERA system from an incoming message.
 * @export
 * @interface InterfaceSetupMessageIndicatorType
 */
export interface InterfaceSetupMessageIndicatorType {
    /**
     * Character Indicator
     * @type {string}
     * @memberof InterfaceSetupMessageIndicatorType
     */
    characterIndicator?: string;
    /**
     * Numeric Indicator
     * @type {number}
     * @memberof InterfaceSetupMessageIndicatorType
     */
    numericIndicator?: number;
    /**
     * Date Indicator.
     * @type {Date}
     * @memberof InterfaceSetupMessageIndicatorType
     */
    dateIndicator?: Date;
}

/**
 * Check if a given object implements the InterfaceSetupMessageIndicatorType interface.
 */
export function instanceOfInterfaceSetupMessageIndicatorType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InterfaceSetupMessageIndicatorTypeFromJSON(json: any): InterfaceSetupMessageIndicatorType {
    return InterfaceSetupMessageIndicatorTypeFromJSONTyped(json, false);
}

export function InterfaceSetupMessageIndicatorTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceSetupMessageIndicatorType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'characterIndicator': !exists(json, 'characterIndicator') ? undefined : json['characterIndicator'],
        'numericIndicator': !exists(json, 'numericIndicator') ? undefined : json['numericIndicator'],
        'dateIndicator': !exists(json, 'dateIndicator') ? undefined : (new Date(json['dateIndicator'])),
    };
}

export function InterfaceSetupMessageIndicatorTypeToJSON(value?: InterfaceSetupMessageIndicatorType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'characterIndicator': value.characterIndicator,
        'numericIndicator': value.numericIndicator,
        'dateIndicator': value.dateIndicator === undefined ? undefined : (value.dateIndicator.toISOString().substr(0,10)),
    };
}

