/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ReservationInfoType } from './ReservationInfoType';
import {
    ReservationInfoTypeFromJSON,
    ReservationInfoTypeFromJSONTyped,
    ReservationInfoTypeToJSON,
} from './ReservationInfoType';

/**
 * A collection of reservation history details attached to Profiles.
 * @export
 * @interface HistoryListType
 */
export interface HistoryListType {
    /**
     * Additional reservation information attached to the profile . Eg : History reservation details
     * @type {Array<ReservationInfoType>}
     * @memberof HistoryListType
     */
    reservationInfo?: Array<ReservationInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof HistoryListType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof HistoryListType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof HistoryListType
     */
    count?: number;
}

/**
 * Check if a given object implements the HistoryListType interface.
 */
export function instanceOfHistoryListType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HistoryListTypeFromJSON(json: any): HistoryListType {
    return HistoryListTypeFromJSONTyped(json, false);
}

export function HistoryListTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HistoryListType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationInfo': !exists(json, 'reservationInfo') ? undefined : ((json['reservationInfo'] as Array<any>).map(ReservationInfoTypeFromJSON)),
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function HistoryListTypeToJSON(value?: HistoryListType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationInfo': value.reservationInfo === undefined ? undefined : ((value.reservationInfo as Array<any>).map(ReservationInfoTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

