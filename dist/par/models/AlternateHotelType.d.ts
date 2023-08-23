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
import type { AddressType } from './AddressType';
import type { ChannelSummaryInfoType } from './ChannelSummaryInfoType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { CurrencyExchangeRateType } from './CurrencyExchangeRateType';
import type { HotelAmenityType } from './HotelAmenityType';
import type { RateRoomDetailsType } from './RateRoomDetailsType';
import type { RelativePositionType } from './RelativePositionType';
import type { SellMessagesType } from './SellMessagesType';
import type { TelephoneType } from './TelephoneType';
/**
 *
 * @export
 * @interface AlternateHotelType
 */
export interface AlternateHotelType {
    /**
     *
     * @type {AddressType}
     * @memberof AlternateHotelType
     */
    address?: AddressType;
    /**
     *
     * @type {Array<TelephoneType>}
     * @memberof AlternateHotelType
     */
    contactNumbers?: Array<TelephoneType>;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof AlternateHotelType
     */
    minRate?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof AlternateHotelType
     */
    maxRate?: CurrencyAmountType;
    /**
     *
     * @type {RelativePositionType}
     * @memberof AlternateHotelType
     */
    relativePosition?: RelativePositionType;
    /**
     *
     * @type {Array<HotelAmenityType>}
     * @memberof AlternateHotelType
     */
    hotelAmenities?: Array<HotelAmenityType>;
    /**
     * True if the hotel has any scheduled events within the requested date range.
     * @type {boolean}
     * @memberof AlternateHotelType
     */
    event?: boolean;
    /**
     *
     * @type {RateRoomDetailsType}
     * @memberof AlternateHotelType
     */
    rateRoomDetails?: RateRoomDetailsType;
    /**
     *
     * @type {ChannelSummaryInfoType}
     * @memberof AlternateHotelType
     */
    channelSummaryInfo?: ChannelSummaryInfoType;
    /**
     *
     * @type {SellMessagesType}
     * @memberof AlternateHotelType
     */
    sellMessages?: SellMessagesType;
    /**
     * Exchange Rate information for a currency code.
     * @type {Array<CurrencyExchangeRateType>}
     * @memberof AlternateHotelType
     */
    currencyExchangeRates?: Array<CurrencyExchangeRateType>;
    /**
     * The code that identifies a hotel chain or management group. The hotel chain code is decided between vendors. This attribute is optional if the hotel is an independent property that can be identified by the HotelCode attribute.
     * @type {string}
     * @memberof AlternateHotelType
     */
    chainCode?: string;
    /**
     * The code that uniquely identifies a single hotel property. The hotel code is decided between vendors.
     * @type {string}
     * @memberof AlternateHotelType
     */
    hotelId?: string;
    /**
     * The IATA city code; for example DCA, ORD.
     * @type {string}
     * @memberof AlternateHotelType
     */
    hotelCityCode?: string;
    /**
     * A text field used to communicate the proper name of the hotel.
     * @type {string}
     * @memberof AlternateHotelType
     */
    hotelName?: string;
    /**
     * A text field used to communicate the context (or source of - ex Sabre, Galileo, Worldspan, Amadeus) the HotelReferenceGroup codes.
     * @type {string}
     * @memberof AlternateHotelType
     */
    hotelCodeContext?: string;
    /**
     * The name of the hotel chain (e.g., Hilton, Marriott, Hyatt).
     * @type {string}
     * @memberof AlternateHotelType
     */
    chainName?: string;
    /**
     *
     * @type {string}
     * @memberof AlternateHotelType
     */
    hotelType?: string;
    /**
     *
     * @type {boolean}
     * @memberof AlternateHotelType
     */
    negotiated?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof AlternateHotelType
     */
    alternate?: boolean;
    /**
     *
     * @type {number}
     * @memberof AlternateHotelType
     */
    displayOrder?: number;
    /**
     *
     * @type {number}
     * @memberof AlternateHotelType
     */
    inventoryRooms?: number;
    /**
     *
     * @type {number}
     * @memberof AlternateHotelType
     */
    availableRooms?: number;
    /**
     * Comments of alternate Hotel.
     * @type {string}
     * @memberof AlternateHotelType
     */
    comments?: string;
    /**
     * Compass direction to the attraction from the hotel (North/South, etc).
     * @type {string}
     * @memberof AlternateHotelType
     */
    direction?: string;
    /**
     * whether the alternate relationship should be applied to the alternate hotel as well. If this flag is true, not only will the AlternateHotelCode hotel be an alternate for HotelCode hotel, but HotelCode hotel will also be an alternate for AlternateHotelCode hotel.
     * @type {boolean}
     * @memberof AlternateHotelType
     */
    reciprocalRelationship?: boolean;
    /**
     *
     * @type {string}
     * @memberof AlternateHotelType
     */
    alternateHotelCode?: string;
    /**
     *
     * @type {string}
     * @memberof AlternateHotelType
     */
    newAlternateHotelCode?: string;
}
/**
 * Check if a given object implements the AlternateHotelType interface.
 */
export declare function instanceOfAlternateHotelType(value: object): boolean;
export declare function AlternateHotelTypeFromJSON(json: any): AlternateHotelType;
export declare function AlternateHotelTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AlternateHotelType;
export declare function AlternateHotelTypeToJSON(value?: AlternateHotelType | null): any;
