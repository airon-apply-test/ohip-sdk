/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SellLimitByDateType } from './SellLimitByDateType';
import {
    SellLimitByDateTypeFromJSON,
    SellLimitByDateTypeFromJSONTyped,
    SellLimitByDateTypeToJSON,
} from './SellLimitByDateType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * A Request message that sets sell limit for a particular date for all the levels House, room type and room class. The request can contain the number type in which the sell limits is to be fetched for the different levels.
 * @export
 * @interface SellLimit
 */
export interface SellLimit {
    /**
     * 
     * @type {SellLimitByDateType}
     * @memberof SellLimit
     */
    sellLimitsByDate?: SellLimitByDateType;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof SellLimit
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the SellLimit interface.
 */
export function instanceOfSellLimit(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SellLimitFromJSON(json: any): SellLimit {
    return SellLimitFromJSONTyped(json, false);
}

export function SellLimitFromJSONTyped(json: any, ignoreDiscriminator: boolean): SellLimit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sellLimitsByDate': !exists(json, 'sellLimitsByDate') ? undefined : SellLimitByDateTypeFromJSON(json['sellLimitsByDate']),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function SellLimitToJSON(value?: SellLimit | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sellLimitsByDate': SellLimitByDateTypeToJSON(value.sellLimitsByDate),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

