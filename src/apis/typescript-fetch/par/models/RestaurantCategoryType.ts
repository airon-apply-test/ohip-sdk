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


/**
 * Simple type for Hotel Restaurant.
 * @export
 */
export const RestaurantCategoryType = {
    Restaurant: 'Restaurant',
    Retail: 'Retail'
} as const;
export type RestaurantCategoryType = typeof RestaurantCategoryType[keyof typeof RestaurantCategoryType];


export function RestaurantCategoryTypeFromJSON(json: any): RestaurantCategoryType {
    return RestaurantCategoryTypeFromJSONTyped(json, false);
}

export function RestaurantCategoryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RestaurantCategoryType {
    return json as RestaurantCategoryType;
}

export function RestaurantCategoryTypeToJSON(value?: RestaurantCategoryType | null): any {
    return value as any;
}

