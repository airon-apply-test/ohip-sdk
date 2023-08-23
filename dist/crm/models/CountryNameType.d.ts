/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Code for a country or a nationality.
 * @export
 * @interface CountryNameType
 */
export interface CountryNameType {
    /**
     * Used for Character Strings, length 0 to 200.
     * @type {string}
     * @memberof CountryNameType
     */
    value?: string;
    /**
     * Code for a country or a nationality.
     * @type {string}
     * @memberof CountryNameType
     */
    code?: string;
}
/**
 * Check if a given object implements the CountryNameType interface.
 */
export declare function instanceOfCountryNameType(value: object): boolean;
export declare function CountryNameTypeFromJSON(json: any): CountryNameType;
export declare function CountryNameTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CountryNameType;
export declare function CountryNameTypeToJSON(value?: CountryNameType | null): any;
