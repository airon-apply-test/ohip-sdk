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
import type { StatisticSetType } from './StatisticSetType';
import {
    StatisticSetTypeFromJSON,
    StatisticSetTypeFromJSONTyped,
    StatisticSetTypeToJSON,
} from './StatisticSetType';

/**
 * Defines the codes and corresponding categories for which the data in the other elements has been gathered.
 * @export
 * @interface StatisticCodeType
 */
export interface StatisticCodeType {
    /**
     * Collection of statistic summary data.
     * @type {Array<StatisticSetType>}
     * @memberof StatisticCodeType
     */
    statisticDate?: Array<StatisticSetType>;
    /**
     * Actual code used by the system to collect the statistics (e.g. CORP, RACK if category is Market Segment).
     * @type {string}
     * @memberof StatisticCodeType
     */
    statCode?: string;
    /**
     * Category Code category of StatCode attribute (e.g. Market Segment).
     * @type {string}
     * @memberof StatisticCodeType
     */
    statCategoryCode?: string;
    /**
     * Class grouping of StatCode attribute.
     * @type {string}
     * @memberof StatisticCodeType
     */
    statCodeClass?: string;
    /**
     * Statistic code description.
     * @type {string}
     * @memberof StatisticCodeType
     */
    description?: string;
}

/**
 * Check if a given object implements the StatisticCodeType interface.
 */
export function instanceOfStatisticCodeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StatisticCodeTypeFromJSON(json: any): StatisticCodeType {
    return StatisticCodeTypeFromJSONTyped(json, false);
}

export function StatisticCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StatisticCodeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'statisticDate': !exists(json, 'statisticDate') ? undefined : ((json['statisticDate'] as Array<any>).map(StatisticSetTypeFromJSON)),
        'statCode': !exists(json, 'statCode') ? undefined : json['statCode'],
        'statCategoryCode': !exists(json, 'statCategoryCode') ? undefined : json['statCategoryCode'],
        'statCodeClass': !exists(json, 'statCodeClass') ? undefined : json['statCodeClass'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function StatisticCodeTypeToJSON(value?: StatisticCodeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'statisticDate': value.statisticDate === undefined ? undefined : ((value.statisticDate as Array<any>).map(StatisticSetTypeToJSON)),
        'statCode': value.statCode,
        'statCategoryCode': value.statCategoryCode,
        'statCodeClass': value.statCodeClass,
        'description': value.description,
    };
}

