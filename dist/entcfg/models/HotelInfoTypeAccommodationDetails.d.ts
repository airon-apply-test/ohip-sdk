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
/**
 * Accommodation Details of the hotel.
 * @export
 * @interface HotelInfoTypeAccommodationDetails
 */
export interface HotelInfoTypeAccommodationDetails {
    /**
     * The number of Single Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    singleRooms?: number;
    /**
     * The number of Double Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    doubleRooms?: number;
    /**
     * The number of Twin Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    twinRooms?: number;
    /**
     * The number of Family Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    familyRooms?: number;
    /**
     * The number of Connecting Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    connectingRooms?: number;
    /**
     * The number of Accessible Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    accessibleRooms?: number;
    /**
     * The number of Non-Smoking Rooms of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    nonSmokingRooms?: number;
    /**
     * Maximum Adults for Family Room Type.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    maxAdultsInFamilyRoom?: number;
    /**
     * Maximum Children for Family Room Type.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    maxChildrenInFamilyRoom?: number;
    /**
     * The total number of the Guest Room Floors.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    guestRoomFloors?: number;
    /**
     * The number of Guest Room Elevators.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    guestRoomElevators?: number;
    /**
     * The number of Suites of the Hotel.
     * @type {number}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    suites?: number;
    /**
     * The floor number of Executive Floors of the Hotel.
     * @type {string}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    executiveFloorNo?: string;
    /**
     * The information about the Room Amenities.
     * @type {string}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    roomAmenties?: string;
    /**
     * The Description of the shops in the Hotel.
     * @type {string}
     * @memberof HotelInfoTypeAccommodationDetails
     */
    shopDescription?: string;
}
/**
 * Check if a given object implements the HotelInfoTypeAccommodationDetails interface.
 */
export declare function instanceOfHotelInfoTypeAccommodationDetails(value: object): boolean;
export declare function HotelInfoTypeAccommodationDetailsFromJSON(json: any): HotelInfoTypeAccommodationDetails;
export declare function HotelInfoTypeAccommodationDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInfoTypeAccommodationDetails;
export declare function HotelInfoTypeAccommodationDetailsToJSON(value?: HotelInfoTypeAccommodationDetails | null): any;
