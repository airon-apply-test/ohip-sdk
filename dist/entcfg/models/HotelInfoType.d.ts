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
import type { AlternateHotelType } from './AlternateHotelType';
import type { CommentInfoType } from './CommentInfoType';
import type { HotelAttractionType } from './HotelAttractionType';
import type { HotelContactType } from './HotelContactType';
import type { HotelCorporateInformationsType } from './HotelCorporateInformationsType';
import type { HotelEventSpacesType } from './HotelEventSpacesType';
import type { HotelInfoTypeAccommodationDetails } from './HotelInfoTypeAccommodationDetails';
import type { HotelInfoTypeAddress } from './HotelInfoTypeAddress';
import type { HotelInfoTypeCommunication } from './HotelInfoTypeCommunication';
import type { HotelInfoTypeGeneralInformation } from './HotelInfoTypeGeneralInformation';
import type { HotelInfoTypePrimaryDetails } from './HotelInfoTypePrimaryDetails';
import type { HotelInfoTypePropertyControls } from './HotelInfoTypePropertyControls';
import type { HotelRateRangeType } from './HotelRateRangeType';
import type { HotelRestaurantType } from './HotelRestaurantType';
import type { MeetingRoomType } from './MeetingRoomType';
/**
 * Contains the basic configuration information about a Hotel.
 * @export
 * @interface HotelInfoType
 */
export interface HotelInfoType {
    /**
     *
     * @type {HotelInfoTypePrimaryDetails}
     * @memberof HotelInfoType
     */
    primaryDetails?: HotelInfoTypePrimaryDetails;
    /**
     *
     * @type {HotelInfoTypeGeneralInformation}
     * @memberof HotelInfoType
     */
    generalInformation?: HotelInfoTypeGeneralInformation;
    /**
     *
     * @type {HotelInfoTypeAccommodationDetails}
     * @memberof HotelInfoType
     */
    accommodationDetails?: HotelInfoTypeAccommodationDetails;
    /**
     *
     * @type {HotelInfoTypePropertyControls}
     * @memberof HotelInfoType
     */
    propertyControls?: HotelInfoTypePropertyControls;
    /**
     *
     * @type {HotelInfoTypeCommunication}
     * @memberof HotelInfoType
     */
    communication?: HotelInfoTypeCommunication;
    /**
     *
     * @type {HotelInfoTypeAddress}
     * @memberof HotelInfoType
     */
    address?: HotelInfoTypeAddress;
    /**
     *
     * @type {Array<HotelRestaurantType>}
     * @memberof HotelInfoType
     */
    hotelRestaurants?: Array<HotelRestaurantType>;
    /**
     * Lists of rate ranges of the hotel.
     * @type {Array<HotelRateRangeType>}
     * @memberof HotelInfoType
     */
    hotelRateRanges?: Array<HotelRateRangeType>;
    /**
     *
     * @type {Array<AlternateHotelType>}
     * @memberof HotelInfoType
     */
    alternateHotels?: Array<AlternateHotelType>;
    /**
     * Lists of contacts of the hotel.
     * @type {Array<HotelContactType>}
     * @memberof HotelInfoType
     */
    hotelContacts?: Array<HotelContactType>;
    /**
     *
     * @type {HotelEventSpacesType}
     * @memberof HotelInfoType
     */
    hotelEventSpaces?: HotelEventSpacesType;
    /**
     * List of Notes of the hotel.
     * @type {Array<CommentInfoType>}
     * @memberof HotelInfoType
     */
    hotelNotes?: Array<CommentInfoType>;
    /**
     *
     * @type {HotelCorporateInformationsType}
     * @memberof HotelInfoType
     */
    hotelCorporateInformations?: HotelCorporateInformationsType;
    /**
     *
     * @type {Array<HotelAttractionType>}
     * @memberof HotelInfoType
     */
    attractions?: Array<HotelAttractionType>;
    /**
     * List of meeting rooms of the hotel.
     * @type {Array<MeetingRoomType>}
     * @memberof HotelInfoType
     */
    meetingRooms?: Array<MeetingRoomType>;
    /**
     * The code that identifies a hotel chain or management group. The hotel chain code is decided between vendors. This attribute is optional if the hotel is an independent property that can be identified by the HotelCode attribute.
     * @type {string}
     * @memberof HotelInfoType
     */
    chainCode?: string;
    /**
     * The code that uniquely identifies a single hotel property. The hotel code is decided between vendors.
     * @type {string}
     * @memberof HotelInfoType
     */
    hotelId?: string;
    /**
     * The IATA city code; for example DCA, ORD.
     * @type {string}
     * @memberof HotelInfoType
     */
    hotelCityCode?: string;
    /**
     * A text field used to communicate the proper name of the hotel.
     * @type {string}
     * @memberof HotelInfoType
     */
    hotelName?: string;
    /**
     * A text field used to communicate the context (or source of - ex Sabre, Galileo, Worldspan, Amadeus) the HotelReferenceGroup codes.
     * @type {string}
     * @memberof HotelInfoType
     */
    hotelCodeContext?: string;
    /**
     * The name of the hotel chain (e.g., Hilton, Marriott, Hyatt).
     * @type {string}
     * @memberof HotelInfoType
     */
    chainName?: string;
}
/**
 * Check if a given object implements the HotelInfoType interface.
 */
export declare function instanceOfHotelInfoType(value: object): boolean;
export declare function HotelInfoTypeFromJSON(json: any): HotelInfoType;
export declare function HotelInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInfoType;
export declare function HotelInfoTypeToJSON(value?: HotelInfoType | null): any;
