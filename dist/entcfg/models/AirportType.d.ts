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
import type { TransportationInfoType } from './TransportationInfoType';
/**
 * Base details of the airport.
 * @export
 * @interface AirportType
 */
export interface AirportType {
    /**
     * Description of the airport.
     * @type {string}
     * @memberof AirportType
     */
    description?: string;
    /**
     * Distance from the hotel to the airport.
     * @type {number}
     * @memberof AirportType
     */
    distance?: number;
    /**
     * Unit of distance for the Distance measurement.
     * @type {string}
     * @memberof AirportType
     */
    distanceType?: string;
    /**
     * Driving time from the hotel to the airport.
     * @type {string}
     * @memberof AirportType
     */
    drivingTime?: string;
    /**
     * Direction of the airport in relation to the hotel.
     * @type {string}
     * @memberof AirportType
     */
    direction?: string;
    /**
     * URL of the airport's website.
     * @type {string}
     * @memberof AirportType
     */
    website?: string;
    /**
     * Transportation option available for the airport.
     * @type {Array<TransportationInfoType>}
     * @memberof AirportType
     */
    transportationOptions?: Array<TransportationInfoType>;
    /**
     * Sequence number for displaying the airport.
     * @type {number}
     * @memberof AirportType
     */
    sequence?: number;
    /**
     * Airport code identifying the airport.
     * @type {string}
     * @memberof AirportType
     */
    code?: string;
    /**
     * Hotel code that the airport belongs to.
     * @type {string}
     * @memberof AirportType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the AirportType interface.
 */
export declare function instanceOfAirportType(value: object): boolean;
export declare function AirportTypeFromJSON(json: any): AirportType;
export declare function AirportTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AirportType;
export declare function AirportTypeToJSON(value?: AirportType | null): any;
