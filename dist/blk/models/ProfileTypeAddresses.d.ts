/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
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
 * @interface ProfileTypeAddresses
 */
export interface ProfileTypeAddresses {
    /**
     * Collection of Detailed information on an address for the customer.
     * @type {Array<AddressInfoType>}
     * @memberof ProfileTypeAddresses
     */
    addressInfo?: Array<AddressInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeAddresses
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeAddresses
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeAddresses
     */
    count?: number;
}
/**
 * Check if a given object implements the ProfileTypeAddresses interface.
 */
export declare function instanceOfProfileTypeAddresses(value: object): boolean;
export declare function ProfileTypeAddressesFromJSON(json: any): ProfileTypeAddresses;
export declare function ProfileTypeAddressesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeAddresses;
export declare function ProfileTypeAddressesToJSON(value?: ProfileTypeAddresses | null): any;
