/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { StayReservationInfoType } from './StayReservationInfoType';
import {
    StayReservationInfoTypeFromJSON,
    StayReservationInfoTypeFromJSONTyped,
    StayReservationInfoTypeToJSON,
} from './StayReservationInfoType';

/**
 * A collection of reservation history details attached to Profiles.
 * @export
 * @interface StayHistoryListType
 */
export interface StayHistoryListType {
    /**
     * Additional reservation information attached to the profile . Eg : History reservation details
     * @type {Array<StayReservationInfoType>}
     * @memberof StayHistoryListType
     */
    reservationInfo?: Array<StayReservationInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof StayHistoryListType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof StayHistoryListType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof StayHistoryListType
     */
    count?: number;
}

/**
 * Check if a given object implements the StayHistoryListType interface.
 */
export function instanceOfStayHistoryListType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StayHistoryListTypeFromJSON(json: any): StayHistoryListType {
    return StayHistoryListTypeFromJSONTyped(json, false);
}

export function StayHistoryListTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StayHistoryListType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationInfo': !exists(json, 'reservationInfo') ? undefined : ((json['reservationInfo'] as Array<any>).map(StayReservationInfoTypeFromJSON)),
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function StayHistoryListTypeToJSON(value?: StayHistoryListType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationInfo': value.reservationInfo === undefined ? undefined : ((value.reservationInfo as Array<any>).map(StayReservationInfoTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

