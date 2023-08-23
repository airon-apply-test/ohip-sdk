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
 * Represents Group profile type.
 * @export
 */
export declare const ExchangeProfileTypeType: {
    readonly Guest: "Guest";
    readonly Company: "Company";
    readonly TravelAgent: "TravelAgent";
    readonly Source: "Source";
    readonly Group: "Group";
};
export type ExchangeProfileTypeType = typeof ExchangeProfileTypeType[keyof typeof ExchangeProfileTypeType];
export declare function ExchangeProfileTypeTypeFromJSON(json: any): ExchangeProfileTypeType;
export declare function ExchangeProfileTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExchangeProfileTypeType;
export declare function ExchangeProfileTypeTypeToJSON(value?: ExchangeProfileTypeType | null): any;
