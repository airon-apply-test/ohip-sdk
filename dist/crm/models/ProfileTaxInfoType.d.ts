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
 * Profile information related to tax.
 * @export
 * @interface ProfileTaxInfoType
 */
export interface ProfileTaxInfoType {
    /**
     * The tax id of this profile. Usually issued by a government agency. Used by 1099 printing.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    tax1No?: string;
    /**
     * Tax 2 id of this profile.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    tax2No?: string;
    /**
     * Tax Category to be changed.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    taxCategory?: string;
    /**
     * Tax Office to be changed.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    taxOffice?: string;
    /**
     * Tax type to be changed.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    taxType?: string;
    /**
     * Business ID. The maximum length of this element should not exceed 120 characters.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    businessId?: string;
    /**
     * Business Registration Code. The maximum length of this element should not exceed 120 characters.
     * @type {string}
     * @memberof ProfileTaxInfoType
     */
    businessRegistration?: string;
}
/**
 * Check if a given object implements the ProfileTaxInfoType interface.
 */
export declare function instanceOfProfileTaxInfoType(value: object): boolean;
export declare function ProfileTaxInfoTypeFromJSON(json: any): ProfileTaxInfoType;
export declare function ProfileTaxInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTaxInfoType;
export declare function ProfileTaxInfoTypeToJSON(value?: ProfileTaxInfoType | null): any;
