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
import type { AddressInfoType } from './AddressInfoType';
/**
 * List of customer addresses.
 * @export
 * @interface CompanyProfileTypeAddresses
 */
export interface CompanyProfileTypeAddresses {
    /**
     * Collection of Detailed information on an address for the customer.
     * @type {Array<AddressInfoType>}
     * @memberof CompanyProfileTypeAddresses
     */
    addressInfo?: Array<AddressInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof CompanyProfileTypeAddresses
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof CompanyProfileTypeAddresses
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof CompanyProfileTypeAddresses
     */
    count?: number;
}
/**
 * Check if a given object implements the CompanyProfileTypeAddresses interface.
 */
export declare function instanceOfCompanyProfileTypeAddresses(value: object): boolean;
export declare function CompanyProfileTypeAddressesFromJSON(json: any): CompanyProfileTypeAddresses;
export declare function CompanyProfileTypeAddressesFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompanyProfileTypeAddresses;
export declare function CompanyProfileTypeAddressesToJSON(value?: CompanyProfileTypeAddresses | null): any;
