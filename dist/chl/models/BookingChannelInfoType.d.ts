/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { HotelMappingType } from './HotelMappingType';
import type { ParameterType } from './ParameterType';
/**
 *
 * @export
 * @interface BookingChannelInfoType
 */
export interface BookingChannelInfoType {
    /**
     * Hotel Code mapping for the Booking Channel.
     * @type {Array<HotelMappingType>}
     * @memberof BookingChannelInfoType
     */
    hotelMappings?: Array<HotelMappingType>;
    /**
     * Collection of generic Name-Value-Pair parameters.
     * @type {Array<ParameterType>}
     * @memberof BookingChannelInfoType
     */
    configurationParameters?: Array<ParameterType>;
    /**
     *
     * @type {string}
     * @memberof BookingChannelInfoType
     */
    bookingChannelCode?: string;
    /**
     *
     * @type {string}
     * @memberof BookingChannelInfoType
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof BookingChannelInfoType
     */
    bookingChannelType?: string;
}
/**
 * Check if a given object implements the BookingChannelInfoType interface.
 */
export declare function instanceOfBookingChannelInfoType(value: object): boolean;
export declare function BookingChannelInfoTypeFromJSON(json: any): BookingChannelInfoType;
export declare function BookingChannelInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BookingChannelInfoType;
export declare function BookingChannelInfoTypeToJSON(value?: BookingChannelInfoType | null): any;
