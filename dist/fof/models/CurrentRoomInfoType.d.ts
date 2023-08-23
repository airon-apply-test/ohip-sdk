/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Room information of the reservation for the current day.
 * @export
 * @interface CurrentRoomInfoType
 */
export interface CurrentRoomInfoType {
    /**
     * Current room type.
     * @type {string}
     * @memberof CurrentRoomInfoType
     */
    roomType?: string;
    /**
     * Current room number.
     * @type {string}
     * @memberof CurrentRoomInfoType
     */
    roomId?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof CurrentRoomInfoType
     */
    suggestedRoomNumbers?: Array<string>;
    /**
     * Current room description.
     * @type {string}
     * @memberof CurrentRoomInfoType
     */
    roomDescription?: string;
    /**
     * Represents the room view code like City view, River view, Ocean view etc.
     * @type {string}
     * @memberof CurrentRoomInfoType
     */
    roomViewCode?: string;
    /**
     * Represents the room was assigned by AI Room Assignment.
     * @type {boolean}
     * @memberof CurrentRoomInfoType
     */
    assignedByAI?: boolean;
    /**
     * Represents the room was upgraded by AI Room Assignment.
     * @type {boolean}
     * @memberof CurrentRoomInfoType
     */
    upgradedByAI?: boolean;
}
/**
 * Check if a given object implements the CurrentRoomInfoType interface.
 */
export declare function instanceOfCurrentRoomInfoType(value: object): boolean;
export declare function CurrentRoomInfoTypeFromJSON(json: any): CurrentRoomInfoType;
export declare function CurrentRoomInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CurrentRoomInfoType;
export declare function CurrentRoomInfoTypeToJSON(value?: CurrentRoomInfoType | null): any;
