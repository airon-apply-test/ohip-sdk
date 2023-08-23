/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
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
export declare function instanceOfMealPlanCodeType(value: object): boolean;
export declare function MealPlanCodeTypeFromJSON(json: any): MealPlanCodeType;
export declare function MealPlanCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MealPlanCodeType;
export declare function MealPlanCodeTypeToJSON(value?: MealPlanCodeType | null): any;
