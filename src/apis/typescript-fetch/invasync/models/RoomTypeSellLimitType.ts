/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Inventory Asynchronous API
 * APIs to cater for Inventory related asynchronous functionality in OPERA Cloud. This includes viewing inventory data along with its revenue and updating inventory&apos;s sell limits for date ranges. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
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
 * @interface RoomTypeSellLimitType
 */
export interface RoomTypeSellLimitType {
    /**
     * 
     * @type {Date}
     * @memberof RoomTypeSellLimitType
     */
    date?: Date;
    /**
     * 
     * @type {number}
     * @memberof RoomTypeSellLimitType
     */
    amount?: number;
    /**
     * Indicates if sell limit is flat or percentage.
     * @type {string}
     * @memberof RoomTypeSellLimitType
     */
    flatOrPercentage?: RoomTypeSellLimitTypeFlatOrPercentageEnum;
    /**
     * 
     * @type {string}
     * @memberof RoomTypeSellLimitType
     */
    roomType?: string;
}


/**
 * @export
 */
export const RoomTypeSellLimitTypeFlatOrPercentageEnum = {
    Flat: 'Flat',
    Percentage: 'Percentage'
} as const;
export type RoomTypeSellLimitTypeFlatOrPercentageEnum = typeof RoomTypeSellLimitTypeFlatOrPercentageEnum[keyof typeof RoomTypeSellLimitTypeFlatOrPercentageEnum];


/**
 * Check if a given object implements the RoomTypeSellLimitType interface.
 */
export function instanceOfRoomTypeSellLimitType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomTypeSellLimitTypeFromJSON(json: any): RoomTypeSellLimitType {
    return RoomTypeSellLimitTypeFromJSONTyped(json, false);
}

export function RoomTypeSellLimitTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomTypeSellLimitType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'date') ? undefined : (new Date(json['date'])),
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'flatOrPercentage': !exists(json, 'flatOrPercentage') ? undefined : json['flatOrPercentage'],
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
    };
}

export function RoomTypeSellLimitTypeToJSON(value?: RoomTypeSellLimitType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'date': value.date === undefined ? undefined : (value.date.toISOString().substr(0,10)),
        'amount': value.amount,
        'flatOrPercentage': value.flatOrPercentage,
        'roomType': value.roomType,
    };
}

