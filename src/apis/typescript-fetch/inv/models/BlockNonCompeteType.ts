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
import type { RateProtectionType } from './RateProtectionType';
import {
    RateProtectionTypeFromJSON,
    RateProtectionTypeFromJSONTyped,
    RateProtectionTypeToJSON,
} from './RateProtectionType';

/**
 * Block Non Compete code information.
 * @export
 * @interface BlockNonCompeteType
 */
export interface BlockNonCompeteType {
    /**
     * Specifies the Non-Compete Industry.>
     * @type {string}
     * @memberof BlockNonCompeteType
     */
    industry?: string;
    /**
     * Specifies the Non-Compete Industry description.>
     * @type {string}
     * @memberof BlockNonCompeteType
     */
    industryDescription?: string;
    /**
     * 
     * @type {RateProtectionType}
     * @memberof BlockNonCompeteType
     */
    criteria?: RateProtectionType;
    /**
     * Specifies a single date.
     * @type {Array<Date>}
     * @memberof BlockNonCompeteType
     */
    protectedDates?: Array<Date>;
}

/**
 * Check if a given object implements the BlockNonCompeteType interface.
 */
export function instanceOfBlockNonCompeteType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockNonCompeteTypeFromJSON(json: any): BlockNonCompeteType {
    return BlockNonCompeteTypeFromJSONTyped(json, false);
}

export function BlockNonCompeteTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockNonCompeteType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'industry': !exists(json, 'industry') ? undefined : json['industry'],
        'industryDescription': !exists(json, 'industryDescription') ? undefined : json['industryDescription'],
        'criteria': !exists(json, 'criteria') ? undefined : RateProtectionTypeFromJSON(json['criteria']),
        'protectedDates': !exists(json, 'protectedDates') ? undefined : json['protectedDates'],
    };
}

export function BlockNonCompeteTypeToJSON(value?: BlockNonCompeteType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'industry': value.industry,
        'industryDescription': value.industryDescription,
        'criteria': RateProtectionTypeToJSON(value.criteria),
        'protectedDates': value.protectedDates,
    };
}

