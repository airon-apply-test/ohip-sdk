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
/**
 * Holds percentage values for each day.
 * @export
 * @interface BlockAllocationWashTypeOccPercentByDay
 */
export interface BlockAllocationWashTypeOccPercentByDay {
    /**
     * Percentage occupancy value for Sunday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    sunday?: number;
    /**
     * Percentage occupancy value for Monday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    monday?: number;
    /**
     * Percentage occupancy value for Tuesday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    tuesday?: number;
    /**
     * Percentage occupancy value for Wednesday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    wednesday?: number;
    /**
     * Percentage occupancy value for Thursday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    thursday?: number;
    /**
     * Percentage occupancy value for Friday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    friday?: number;
    /**
     * Percentage occupancy value for Saturday.
     * @type {number}
     * @memberof BlockAllocationWashTypeOccPercentByDay
     */
    saturday?: number;
}

/**
 * Check if a given object implements the BlockAllocationWashTypeOccPercentByDay interface.
 */
export function instanceOfBlockAllocationWashTypeOccPercentByDay(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockAllocationWashTypeOccPercentByDayFromJSON(json: any): BlockAllocationWashTypeOccPercentByDay {
    return BlockAllocationWashTypeOccPercentByDayFromJSONTyped(json, false);
}

export function BlockAllocationWashTypeOccPercentByDayFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockAllocationWashTypeOccPercentByDay {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sunday': !exists(json, 'sunday') ? undefined : json['sunday'],
        'monday': !exists(json, 'monday') ? undefined : json['monday'],
        'tuesday': !exists(json, 'tuesday') ? undefined : json['tuesday'],
        'wednesday': !exists(json, 'wednesday') ? undefined : json['wednesday'],
        'thursday': !exists(json, 'thursday') ? undefined : json['thursday'],
        'friday': !exists(json, 'friday') ? undefined : json['friday'],
        'saturday': !exists(json, 'saturday') ? undefined : json['saturday'],
    };
}

export function BlockAllocationWashTypeOccPercentByDayToJSON(value?: BlockAllocationWashTypeOccPercentByDay | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sunday': value.sunday,
        'monday': value.monday,
        'tuesday': value.tuesday,
        'wednesday': value.wednesday,
        'thursday': value.thursday,
        'friday': value.friday,
        'saturday': value.saturday,
    };
}

