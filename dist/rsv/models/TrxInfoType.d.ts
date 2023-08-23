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
/**
 * Transaction codes info.
 * @export
 * @interface TrxInfoType
 */
export interface TrxInfoType {
    /**
     * Transaction codes info.
     * @type {string}
     * @memberof TrxInfoType
     */
    description?: string;
    /**
     * Contains service type for transaction code.
     * @type {string}
     * @memberof TrxInfoType
     */
    trxServiceType?: string;
    /**
     * Unique identifier for the Transaction code.
     * @type {string}
     * @memberof TrxInfoType
     */
    transactionCode?: string;
    /**
     * Hotel context of the Transaction code.
     * @type {string}
     * @memberof TrxInfoType
     */
    hotelId?: string;
    /**
     * Print receipt flag that tells whether the transaction receipt is to be printed or not. This is based on the transaction code.
     * @type {boolean}
     * @memberof TrxInfoType
     */
    printTrxReceipt?: boolean;
}
/**
 * Check if a given object implements the TrxInfoType interface.
 */
export declare function instanceOfTrxInfoType(value: object): boolean;
export declare function TrxInfoTypeFromJSON(json: any): TrxInfoType;
export declare function TrxInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrxInfoType;
export declare function TrxInfoTypeToJSON(value?: TrxInfoType | null): any;
