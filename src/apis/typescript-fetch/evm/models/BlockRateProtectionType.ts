/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RateProtectionType } from './RateProtectionType';
import {
    RateProtectionTypeFromJSON,
    RateProtectionTypeFromJSONTyped,
    RateProtectionTypeToJSON,
} from './RateProtectionType';

/**
 * Block Rate Protection code information.
 * @export
 * @interface BlockRateProtectionType
 */
export interface BlockRateProtectionType {
    /**
     * 
     * @type {RateProtectionType}
     * @memberof BlockRateProtectionType
     */
    criteria?: RateProtectionType;
    /**
     * Specifies a single date.
     * @type {Array<Date>}
     * @memberof BlockRateProtectionType
     */
    protectedDates?: Array<Date>;
}

/**
 * Check if a given object implements the BlockRateProtectionType interface.
 */
export function instanceOfBlockRateProtectionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockRateProtectionTypeFromJSON(json: any): BlockRateProtectionType {
    return BlockRateProtectionTypeFromJSONTyped(json, false);
}

export function BlockRateProtectionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockRateProtectionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : RateProtectionTypeFromJSON(json['criteria']),
        'protectedDates': !exists(json, 'protectedDates') ? undefined : json['protectedDates'],
    };
}

export function BlockRateProtectionTypeToJSON(value?: BlockRateProtectionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': RateProtectionTypeToJSON(value.criteria),
        'protectedDates': value.protectedDates,
    };
}

