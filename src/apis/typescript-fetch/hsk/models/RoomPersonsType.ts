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
 * @interface RoomPersonsType
 */
export interface RoomPersonsType {
    /**
     * 
     * @type {number}
     * @memberof RoomPersonsType
     */
    frontOfficePersons?: number;
    /**
     * 
     * @type {number}
     * @memberof RoomPersonsType
     */
    houseKeepingPersons?: number;
}

/**
 * Check if a given object implements the RoomPersonsType interface.
 */
export function instanceOfRoomPersonsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomPersonsTypeFromJSON(json: any): RoomPersonsType {
    return RoomPersonsTypeFromJSONTyped(json, false);
}

export function RoomPersonsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomPersonsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'frontOfficePersons': !exists(json, 'frontOfficePersons') ? undefined : json['frontOfficePersons'],
        'houseKeepingPersons': !exists(json, 'houseKeepingPersons') ? undefined : json['houseKeepingPersons'],
    };
}

export function RoomPersonsTypeToJSON(value?: RoomPersonsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'frontOfficePersons': value.frontOfficePersons,
        'houseKeepingPersons': value.houseKeepingPersons,
    };
}

