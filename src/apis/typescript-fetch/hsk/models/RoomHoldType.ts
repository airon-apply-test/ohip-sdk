/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface RoomHoldType
 */
export interface RoomHoldType {
    /**
     * The date and time when hold will expire.
     * @type {Date}
     * @memberof RoomHoldType
     */
    holdUntil?: Date;
    /**
     * User who placed room on hold.
     * @type {string}
     * @memberof RoomHoldType
     */
    holdUser?: string;
    /**
     * Comments of the room hold.
     * @type {string}
     * @memberof RoomHoldType
     */
    comments?: string;
}

/**
 * Check if a given object implements the RoomHoldType interface.
 */
export function instanceOfRoomHoldType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomHoldTypeFromJSON(json: any): RoomHoldType {
    return RoomHoldTypeFromJSONTyped(json, false);
}

export function RoomHoldTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomHoldType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'holdUntil': !exists(json, 'holdUntil') ? undefined : (new Date(json['holdUntil'])),
        'holdUser': !exists(json, 'holdUser') ? undefined : json['holdUser'],
        'comments': !exists(json, 'comments') ? undefined : json['comments'],
    };
}

export function RoomHoldTypeToJSON(value?: RoomHoldType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'holdUntil': value.holdUntil === undefined ? undefined : (value.holdUntil.toISOString()),
        'holdUser': value.holdUser,
        'comments': value.comments,
    };
}

