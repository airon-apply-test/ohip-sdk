/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RatePlanSellControlsType } from './RatePlanSellControlsType';
import {
    RatePlanSellControlsTypeFromJSON,
    RatePlanSellControlsTypeFromJSONTyped,
    RatePlanSellControlsTypeToJSON,
} from './RatePlanSellControlsType';
import type { RatePlanYieldControlsType } from './RatePlanYieldControlsType';
import {
    RatePlanYieldControlsTypeFromJSON,
    RatePlanYieldControlsTypeFromJSONTyped,
    RatePlanYieldControlsTypeToJSON,
} from './RatePlanYieldControlsType';

/**
 * 
 * @export
 * @interface RatePlanControlsType
 */
export interface RatePlanControlsType {
    /**
     * 
     * @type {RatePlanSellControlsType}
     * @memberof RatePlanControlsType
     */
    sell?: RatePlanSellControlsType;
    /**
     * 
     * @type {RatePlanYieldControlsType}
     * @memberof RatePlanControlsType
     */
    _yield?: RatePlanYieldControlsType;
}

/**
 * Check if a given object implements the RatePlanControlsType interface.
 */
export function instanceOfRatePlanControlsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RatePlanControlsTypeFromJSON(json: any): RatePlanControlsType {
    return RatePlanControlsTypeFromJSONTyped(json, false);
}

export function RatePlanControlsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RatePlanControlsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sell': !exists(json, 'sell') ? undefined : RatePlanSellControlsTypeFromJSON(json['sell']),
        '_yield': !exists(json, 'yield') ? undefined : RatePlanYieldControlsTypeFromJSON(json['yield']),
    };
}

export function RatePlanControlsTypeToJSON(value?: RatePlanControlsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sell': RatePlanSellControlsTypeToJSON(value.sell),
        'yield': RatePlanYieldControlsTypeToJSON(value._yield),
    };
}

