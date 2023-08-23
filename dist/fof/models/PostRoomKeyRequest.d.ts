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
 *
 * @export
 * @interface PostRoomKeyRequest
 */
export interface PostRoomKeyRequest {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof PostRoomKeyRequest
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     * A room number to generate a key for. When creating keys using just the room number, the interface may allow only one-shot keys.
     * @type {string}
     * @memberof PostRoomKeyRequest
     */
    roomNumber?: string;
    /**
     *
     * @type {KeyTrackType}
     * @memberof PostRoomKeyRequest
     */
    keyTrack?: KeyTrackType;
    /**
     *
     * @type {string}
     * @memberof PostRoomKeyRequest
     */
    resort?: string;
    /**
     *
     * @type {number}
     * @memberof PostRoomKeyRequest
     */
    noOfKeys?: number;
    /**
     *
     * @type {Date}
     * @memberof PostRoomKeyRequest
     */
    keyValidityStart?: Date;
    /**
     *
     * @type {Date}
     * @memberof PostRoomKeyRequest
     */
    keyValidityEnd?: Date;
    /**
     *
     * @type {string}
     * @memberof PostRoomKeyRequest
     */
    encoderTerminal?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof PostRoomKeyRequest
     */
    encoderId?: UniqueIDType;
    /**
     *
     * @type {RoomKeyType}
     * @memberof PostRoomKeyRequest
     */
    keyType?: RoomKeyType;
    /**
     *
     * @type {KeyCardType}
     * @memberof PostRoomKeyRequest
     */
    keyCardType?: KeyCardType;
    /**
     *
     * @type {string}
     * @memberof PostRoomKeyRequest
     */
    keyCardUId?: string;
    /**
     *
     * @type {string}
     * @memberof PostRoomKeyRequest
     */
    keyOptions?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof PostRoomKeyRequest
     */
    additionalRooms?: Array<string>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostRoomKeyRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostRoomKeyRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostRoomKeyRequest interface.
 */
export declare function instanceOfPostRoomKeyRequest(value: object): boolean;
export declare function PostRoomKeyRequestFromJSON(json: any): PostRoomKeyRequest;
export declare function PostRoomKeyRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostRoomKeyRequest;
export declare function PostRoomKeyRequestToJSON(value?: PostRoomKeyRequest | null): any;
