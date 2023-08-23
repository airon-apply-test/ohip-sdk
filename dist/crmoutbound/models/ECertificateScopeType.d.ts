/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Indicates that OPERA E-Certificate is available for a specific list of hotels.
 * @export
 */
export declare const ECertificateScopeType: {
    readonly Global: "Global";
    readonly Hotel: "Hotel";
    readonly MultiHotel: "MultiHotel";
};
export type ECertificateScopeType = typeof ECertificateScopeType[keyof typeof ECertificateScopeType];
export declare function ECertificateScopeTypeFromJSON(json: any): ECertificateScopeType;
export declare function ECertificateScopeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateScopeType;
export declare function ECertificateScopeTypeToJSON(value?: ECertificateScopeType | null): any;
