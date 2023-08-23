/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
/**
 * Information about the meeting rooms of the hotel.
 * @export
 * @interface MeetingRoomType
 */
export interface MeetingRoomType {
    /**
     * the code of the meeting room
     * @type {string}
     * @memberof MeetingRoomType
     */
    code?: string;
    /**
     * The charge for the meeting room.
     * @type {string}
     * @memberof MeetingRoomType
     */
    charge?: string;
    /**
     * A description of the meeting room.
     * @type {string}
     * @memberof MeetingRoomType
     */
    description?: string;
}

/**
 * Check if a given object implements the MeetingRoomType interface.
 */
export function instanceOfMeetingRoomType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MeetingRoomTypeFromJSON(json: any): MeetingRoomType {
    return MeetingRoomTypeFromJSONTyped(json, false);
}

export function MeetingRoomTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MeetingRoomType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'charge': !exists(json, 'charge') ? undefined : json['charge'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function MeetingRoomTypeToJSON(value?: MeetingRoomType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'charge': value.charge,
        'description': value.description,
    };
}

