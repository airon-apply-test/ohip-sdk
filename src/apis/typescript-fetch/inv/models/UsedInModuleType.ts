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


/**
 * Simple type for indicating the module for which user wants to create the SellMessage. Valid module types are Look To Book Sales,Reservations,Blocks,Function Diary and Group Rooms Control.
 * @export
 */
export const UsedInModuleType = {
    Blocks: 'Blocks',
    Reservations: 'Reservations',
    FunctionDiary: 'FunctionDiary',
    LookToBookSales: 'LookToBookSales',
    GroupRoomsControl: 'GroupRoomsControl'
} as const;
export type UsedInModuleType = typeof UsedInModuleType[keyof typeof UsedInModuleType];


export function UsedInModuleTypeFromJSON(json: any): UsedInModuleType {
    return UsedInModuleTypeFromJSONTyped(json, false);
}

export function UsedInModuleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsedInModuleType {
    return json as UsedInModuleType;
}

export function UsedInModuleTypeToJSON(value?: UsedInModuleType | null): any {
    return value as any;
}

