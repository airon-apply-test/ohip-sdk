/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RoomTypeSummaryType } from './RoomTypeSummaryType';
import {
    RoomTypeSummaryTypeFromJSON,
    RoomTypeSummaryTypeFromJSONTyped,
    RoomTypeSummaryTypeToJSON,
} from './RoomTypeSummaryType';

/**
 * This type holds collection of room type.
 * @export
 * @interface RoomTypesSummaryType
 */
export interface RoomTypesSummaryType {
    /**
     * 
     * @type {Array<RoomTypeSummaryType>}
     * @memberof RoomTypesSummaryType
     */
    roomTypeSummary?: Array<RoomTypeSummaryType>;
    /**
     * Hotel Code.
     * @type {string}
     * @memberof RoomTypesSummaryType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the RoomTypesSummaryType interface.
 */
export function instanceOfRoomTypesSummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomTypesSummaryTypeFromJSON(json: any): RoomTypesSummaryType {
    return RoomTypesSummaryTypeFromJSONTyped(json, false);
}

export function RoomTypesSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomTypesSummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomTypeSummary': !exists(json, 'roomTypeSummary') ? undefined : ((json['roomTypeSummary'] as Array<any>).map(RoomTypeSummaryTypeFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function RoomTypesSummaryTypeToJSON(value?: RoomTypesSummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomTypeSummary': value.roomTypeSummary === undefined ? undefined : ((value.roomTypeSummary as Array<any>).map(RoomTypeSummaryTypeToJSON)),
        'hotelId': value.hotelId,
    };
}

