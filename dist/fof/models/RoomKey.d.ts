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
import type { InstanceLink } from './InstanceLink';
import type { KeyCardType } from './KeyCardType';
import type { KeyTrackType } from './KeyTrackType';
import type { RoomKeyType } from './RoomKeyType';
import type { UniqueIDType } from './UniqueIDType';
import type { WarningType } from './WarningType';
/**
 * Request for generation of room key.
 * @export
 * @interface RoomKey
 */
export interface RoomKey {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RoomKey
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     * A room number to generate a key for. When creating keys using just the room number, the interface may allow only one-shot keys.
     * @type {string}
     * @memberof RoomKey
     */
    roomNumber?: string;
    /**
     *
     * @type {KeyTrackType}
     * @memberof RoomKey
     */
    keyTrack?: KeyTrackType;
    /**
     *
     * @type {string}
     * @memberof RoomKey
     */
    resort?: string;
    /**
     *
     * @type {number}
     * @memberof RoomKey
     */
    noOfKeys?: number;
    /**
     *
     * @type {Date}
     * @memberof RoomKey
     */
    keyValidityStart?: Date;
    /**
     *
     * @type {Date}
     * @memberof RoomKey
     */
    keyValidityEnd?: Date;
    /**
     *
     * @type {string}
     * @memberof RoomKey
     */
    encoderTerminal?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof RoomKey
     */
    encoderId?: UniqueIDType;
    /**
     *
     * @type {RoomKeyType}
     * @memberof RoomKey
     */
    keyType?: RoomKeyType;
    /**
     *
     * @type {KeyCardType}
     * @memberof RoomKey
     */
    keyCardType?: KeyCardType;
    /**
     *
     * @type {string}
     * @memberof RoomKey
     */
    keyCardUId?: string;
    /**
     *
     * @type {string}
     * @memberof RoomKey
     */
    keyOptions?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof RoomKey
     */
    additionalRooms?: Array<string>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof RoomKey
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RoomKey
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the RoomKey interface.
 */
export declare function instanceOfRoomKey(value: object): boolean;
export declare function RoomKeyFromJSON(json: any): RoomKey;
export declare function RoomKeyFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomKey;
export declare function RoomKeyToJSON(value?: RoomKey | null): any;
