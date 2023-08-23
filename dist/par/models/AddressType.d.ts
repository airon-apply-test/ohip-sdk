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
import type { CountryNameType } from './CountryNameType';
/**
 * Provides address information.
 * @export
 * @interface AddressType
 */
export interface AddressType {
    /**
     * Indicator to define if the Address is validated by the Address Validation System.
     * @type {boolean}
     * @memberof AddressType
     */
    isValidated?: boolean;
    /**
     * When the address is unformatted (FormattedInd="false") these lines will contain free form address details. When the address is formatted and street number and street name must be sent independently, the street number will be sent using StreetNmbr, and the street name will be sent in the first AddressLine occurrence.
     * @type {Array<string>}
     * @memberof AddressType
     */
    addressLine?: Array<string>;
    /**
     * City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
     * @type {string}
     * @memberof AddressType
     */
    cityName?: string;
    /**
     * Post Office Code number.
     * @type {string}
     * @memberof AddressType
     */
    postalCode?: string;
    /**
     * Post Office City Extension Code number. City Extension mainly used for UK addresses.
     * @type {string}
     * @memberof AddressType
     */
    cityExtension?: string;
    /**
     * County or District Name (e.g., Fairfax). This is read only.
     * @type {string}
     * @memberof AddressType
     */
    county?: string;
    /**
     * State or Province name (e.g., Texas).
     * @type {string}
     * @memberof AddressType
     */
    state?: string;
    /**
     *
     * @type {CountryNameType}
     * @memberof AddressType
     */
    country?: CountryNameType;
    /**
     * Language identification.
     * @type {string}
     * @memberof AddressType
     */
    language?: string;
    /**
     * Defines the type of address (e.g. home, business, other).
     * @type {string}
     * @memberof AddressType
     */
    type?: string;
    /**
     * Describes the type code
     * @type {string}
     * @memberof AddressType
     */
    typeDescription?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof AddressType
     */
    primaryInd?: boolean;
    /**
     * Indicates whether to update the reservations or not. If true and the address is primary, then all associated active reservations will be updated with the new primary address.
     * @type {boolean}
     * @memberof AddressType
     */
    updateReservations?: boolean;
    /**
     * The postal barcode for the address.
     * @type {string}
     * @memberof AddressType
     */
    barCode?: string;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof AddressType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof AddressType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof AddressType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof AddressType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof AddressType
     */
    purgeDate?: Date;
}
/**
 * Check if a given object implements the AddressType interface.
 */
export declare function instanceOfAddressType(value: object): boolean;
export declare function AddressTypeFromJSON(json: any): AddressType;
export declare function AddressTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddressType;
export declare function AddressTypeToJSON(value?: AddressType | null): any;
