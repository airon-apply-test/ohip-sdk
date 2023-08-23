/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Meal plan codes associated with the rate codes.
 * @export
 * @interface MealPlanCodeType
 */
export interface MealPlanCodeType {
    /**
     * Meal plan code.
     * @type {string}
     * @memberof MealPlanCodeType
     */
    code?: string;
    /**
     * Represents if the meal plan code is included in rate code or not.
     * @type {boolean}
     * @memberof MealPlanCodeType
     */
    included?: boolean;
    /**
     * Represents if the meal plan code is available for breakfast or not.
     * @type {boolean}
     * @memberof MealPlanCodeType
     */
    breakfastIncluded?: boolean;
    /**
     * Represents if the meal plan code is available for lunch or not.
     * @type {boolean}
     * @memberof MealPlanCodeType
     */
    lunchIncluded?: boolean;
    /**
     * Represents if the meal plan code is available for dinner or not.
     * @type {boolean}
     * @memberof MealPlanCodeType
     */
    dinnerIncluded?: boolean;
}

/**
 * Check if a given object implements the MealPlanCodeType interface.
 */
export function instanceOfMealPlanCodeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MealPlanCodeTypeFromJSON(json: any): MealPlanCodeType {
    return MealPlanCodeTypeFromJSONTyped(json, false);
}

export function MealPlanCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MealPlanCodeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'included': !exists(json, 'included') ? undefined : json['included'],
        'breakfastIncluded': !exists(json, 'breakfastIncluded') ? undefined : json['breakfastIncluded'],
        'lunchIncluded': !exists(json, 'lunchIncluded') ? undefined : json['lunchIncluded'],
        'dinnerIncluded': !exists(json, 'dinnerIncluded') ? undefined : json['dinnerIncluded'],
    };
}

export function MealPlanCodeTypeToJSON(value?: MealPlanCodeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'included': value.included,
        'breakfastIncluded': value.breakfastIncluded,
        'lunchIncluded': value.lunchIncluded,
        'dinnerIncluded': value.dinnerIncluded,
    };
}

