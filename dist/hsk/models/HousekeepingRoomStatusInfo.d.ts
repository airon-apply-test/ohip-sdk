/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Housekeeping Room Status Info Type
 * @export
 * @interface HousekeepingRoomStatusInfo
 */
export interface HousekeepingRoomStatusInfo {
    /**
     *
     * @type {string}
     * @memberof HousekeepingRoomStatusInfo
     */
    housekeepingRoomStatus?: string;
    /**
     *
     * @type {number}
     * @memberof HousekeepingRoomStatusInfo
     */
    numberOfRooms?: number;
    /**
     *
     * @type {string}
     * @memberof HousekeepingRoomStatusInfo
     */
    code?: string;
}
/**
 * Check if a given object implements the HousekeepingRoomStatusInfo interface.
 */
export declare function instanceOfHousekeepingRoomStatusInfo(value: object): boolean;
export declare function HousekeepingRoomStatusInfoFromJSON(json: any): HousekeepingRoomStatusInfo;
export declare function HousekeepingRoomStatusInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): HousekeepingRoomStatusInfo;
export declare function HousekeepingRoomStatusInfoToJSON(value?: HousekeepingRoomStatusInfo | null): any;
