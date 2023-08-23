/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RelativePositionType } from './RelativePositionType';
import type { TelephoneType } from './TelephoneType';
import type { URLType } from './URLType';
/**
 * Information about the transportations of the hotel.
 * @export
 * @interface TransportationType
 */
export interface TransportationType {
    /**
     *
     * @type {TelephoneType}
     * @memberof TransportationType
     */
    phoneNumber?: TelephoneType;
    /**
     *
     * @type {RelativePositionType}
     * @memberof TransportationType
     */
    relativePosition?: RelativePositionType;
    /**
     * The description of the transportation.
     * @type {string}
     * @memberof TransportationType
     */
    description?: string;
    /**
     * Comments about the transportation.
     * @type {string}
     * @memberof TransportationType
     */
    comments?: string;
    /**
     * The price range of the transportation.
     * @type {string}
     * @memberof TransportationType
     */
    priceRange?: string;
    /**
     *
     * @type {URLType}
     * @memberof TransportationType
     */
    website?: URLType;
    /**
     *
     * @type {Array<string>}
     * @memberof TransportationType
     */
    keyOptions?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof TransportationType
     */
    transportationCode?: string;
    /**
     *
     * @type {string}
     * @memberof TransportationType
     */
    label?: string;
    /**
     *
     * @type {number}
     * @memberof TransportationType
     */
    orderBy?: number;
    /**
     * Hotel code for the transportation.
     * @type {string}
     * @memberof TransportationType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the TransportationType interface.
 */
export declare function instanceOfTransportationType(value: object): boolean;
export declare function TransportationTypeFromJSON(json: any): TransportationType;
export declare function TransportationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransportationType;
export declare function TransportationTypeToJSON(value?: TransportationType | null): any;
