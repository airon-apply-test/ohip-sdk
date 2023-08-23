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
 *
 * @export
 */
export declare const ApplicationSettingLevelType: {
    readonly Global: "Global";
    readonly Hotel: "Hotel";
    readonly Cro: "Cro";
};
export type ApplicationSettingLevelType = typeof ApplicationSettingLevelType[keyof typeof ApplicationSettingLevelType];
export declare function ApplicationSettingLevelTypeFromJSON(json: any): ApplicationSettingLevelType;
export declare function ApplicationSettingLevelTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApplicationSettingLevelType;
export declare function ApplicationSettingLevelTypeToJSON(value?: ApplicationSettingLevelType | null): any;
