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
import type { RestrictionSetType } from './RestrictionSetType';
/**
 * Restriction by date range for the requested Hotel.
 * @export
 * @interface RestrictionsByDateRangeType
 */
export interface RestrictionsByDateRangeType {
    /**
     * Restriction set for a date range.
     * @type {Array<RestrictionSetType>}
     * @memberof RestrictionsByDateRangeType
     */
    restrictionSets?: Array<RestrictionSetType>;
    /**
     * The code that identifies a hotel chain or management group. The hotel chain code is decided between vendors. This attribute is optional if the hotel is an independent property that can be identified by the HotelCode attribute.
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    chainCode?: string;
    /**
     * The code that uniquely identifies a single hotel property. The hotel code is decided between vendors.
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    hotelId?: string;
    /**
     * The IATA city code; for example DCA, ORD.
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    hotelCityCode?: string;
    /**
     * A text field used to communicate the proper name of the hotel.
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    hotelName?: string;
    /**
     * A text field used to communicate the context (or source of - ex Sabre, Galileo, Worldspan, Amadeus) the HotelReferenceGroup codes.
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    hotelCodeContext?: string;
    /**
     * The name of the hotel chain (e.g., Hilton, Marriott, Hyatt).
     * @type {string}
     * @memberof RestrictionsByDateRangeType
     */
    chainName?: string;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof RestrictionsByDateRangeType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof RestrictionsByDateRangeType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof RestrictionsByDateRangeType
     */
    count?: number;
}
/**
 * Check if a given object implements the RestrictionsByDateRangeType interface.
 */
export declare function instanceOfRestrictionsByDateRangeType(value: object): boolean;
export declare function RestrictionsByDateRangeTypeFromJSON(json: any): RestrictionsByDateRangeType;
export declare function RestrictionsByDateRangeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RestrictionsByDateRangeType;
export declare function RestrictionsByDateRangeTypeToJSON(value?: RestrictionsByDateRangeType | null): any;
