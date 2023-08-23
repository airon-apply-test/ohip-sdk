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
import type { TotalType } from './TotalType';
import {
    TotalTypeFromJSON,
    TotalTypeFromJSONTyped,
    TotalTypeToJSON,
} from './TotalType';

/**
 * Rate Range details like maximum rate amount and minimum rate amount in each available rate category.
 * @export
 * @interface RateRangeType
 */
export interface RateRangeType {
    /**
     * The base amount charged for the accommodation or service.
     * @type {Array<TotalType>}
     * @memberof RateRangeType
     */
    base?: Array<TotalType>;
    /**
     * Rate Change Indicator.
     * @type {boolean}
     * @memberof RateRangeType
     */
    rateChange?: boolean;
}

/**
 * Check if a given object implements the RateRangeType interface.
 */
export function instanceOfRateRangeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RateRangeTypeFromJSON(json: any): RateRangeType {
    return RateRangeTypeFromJSONTyped(json, false);
}

export function RateRangeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RateRangeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'base': !exists(json, 'base') ? undefined : ((json['base'] as Array<any>).map(TotalTypeFromJSON)),
        'rateChange': !exists(json, 'rateChange') ? undefined : json['rateChange'],
    };
}

export function RateRangeTypeToJSON(value?: RateRangeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'base': value.base === undefined ? undefined : ((value.base as Array<any>).map(TotalTypeToJSON)),
        'rateChange': value.rateChange,
    };
}

