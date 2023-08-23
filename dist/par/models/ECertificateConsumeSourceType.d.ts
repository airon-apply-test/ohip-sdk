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
 * Indicates that OPERA E-Certificate is consumed by hotelPMS.
 * @export
 */
export declare const ECertificateConsumeSourceType: {
    readonly Central: "Central";
    readonly Web: "Web";
    readonly Hotel: "Hotel";
};
export type ECertificateConsumeSourceType = typeof ECertificateConsumeSourceType[keyof typeof ECertificateConsumeSourceType];
export declare function ECertificateConsumeSourceTypeFromJSON(json: any): ECertificateConsumeSourceType;
export declare function ECertificateConsumeSourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateConsumeSourceType;
export declare function ECertificateConsumeSourceTypeToJSON(value?: ECertificateConsumeSourceType | null): any;
