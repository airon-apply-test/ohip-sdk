/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Data modules that are currently supported
 * @export
 */
export const AutomaticTransmissionScheduleModuleTypeType = {
    ArExport: 'ARExport',
    Block: 'Block',
    Inventory: 'Inventory',
    ProfileRequest: 'ProfileRequest',
    ProfileDailyStats: 'ProfileDailyStats'
} as const;
export type AutomaticTransmissionScheduleModuleTypeType = typeof AutomaticTransmissionScheduleModuleTypeType[keyof typeof AutomaticTransmissionScheduleModuleTypeType];


export function AutomaticTransmissionScheduleModuleTypeTypeFromJSON(json: any): AutomaticTransmissionScheduleModuleTypeType {
    return AutomaticTransmissionScheduleModuleTypeTypeFromJSONTyped(json, false);
}

export function AutomaticTransmissionScheduleModuleTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AutomaticTransmissionScheduleModuleTypeType {
    return json as AutomaticTransmissionScheduleModuleTypeType;
}

export function AutomaticTransmissionScheduleModuleTypeTypeToJSON(value?: AutomaticTransmissionScheduleModuleTypeType | null): any {
    return value as any;
}

