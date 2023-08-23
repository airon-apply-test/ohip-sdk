/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
