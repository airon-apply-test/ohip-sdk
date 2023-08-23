/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * The summary info of hotel event space
 * @export
 * @interface HotelEventSpaceSummaryType
 */
export interface HotelEventSpaceSummaryType {
    /**
     * The total event space number which has the same space type, max capacity and setup style.
     * @type {number}
     * @memberof HotelEventSpaceSummaryType
     */
    no?: number;
    /**
     * The type of the event space
     * @type {string}
     * @memberof HotelEventSpaceSummaryType
     */
    spaceType?: string;
    /**
     * The max capacity of this event space group
     * @type {number}
     * @memberof HotelEventSpaceSummaryType
     */
    maxCapacity?: number;
    /**
     * List of event space max occupancy.
     * @type {Array<number>}
     * @memberof HotelEventSpaceSummaryType
     */
    maxOccupancies?: Array<number>;
}

/**
 * Check if a given object implements the HotelEventSpaceSummaryType interface.
 */
export function instanceOfHotelEventSpaceSummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HotelEventSpaceSummaryTypeFromJSON(json: any): HotelEventSpaceSummaryType {
    return HotelEventSpaceSummaryTypeFromJSONTyped(json, false);
}

export function HotelEventSpaceSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelEventSpaceSummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'no': !exists(json, 'no') ? undefined : json['no'],
        'spaceType': !exists(json, 'spaceType') ? undefined : json['spaceType'],
        'maxCapacity': !exists(json, 'maxCapacity') ? undefined : json['maxCapacity'],
        'maxOccupancies': !exists(json, 'maxOccupancies') ? undefined : json['maxOccupancies'],
    };
}

export function HotelEventSpaceSummaryTypeToJSON(value?: HotelEventSpaceSummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'no': value.no,
        'spaceType': value.spaceType,
        'maxCapacity': value.maxCapacity,
        'maxOccupancies': value.maxOccupancies,
    };
}

