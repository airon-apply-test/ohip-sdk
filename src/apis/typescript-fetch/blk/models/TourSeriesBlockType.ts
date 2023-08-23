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
 * Information about a tour series block.
 * @export
 * @interface TourSeriesBlockType
 */
export interface TourSeriesBlockType {
    /**
     * The hotel code of the tour series block.
     * @type {string}
     * @memberof TourSeriesBlockType
     */
    hotelId?: string;
    /**
     * The block code of the tour series block.
     * @type {string}
     * @memberof TourSeriesBlockType
     */
    blockCode?: string;
    /**
     * The start date of the tour series block.
     * @type {Date}
     * @memberof TourSeriesBlockType
     */
    startDate?: Date;
    /**
     * The booking status of the tour series block.
     * @type {string}
     * @memberof TourSeriesBlockType
     */
    blockStatus?: string;
    /**
     * The block name of the tour series block.
     * @type {string}
     * @memberof TourSeriesBlockType
     */
    blockName?: string;
}

/**
 * Check if a given object implements the TourSeriesBlockType interface.
 */
export function instanceOfTourSeriesBlockType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TourSeriesBlockTypeFromJSON(json: any): TourSeriesBlockType {
    return TourSeriesBlockTypeFromJSONTyped(json, false);
}

export function TourSeriesBlockTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TourSeriesBlockType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'blockCode': !exists(json, 'blockCode') ? undefined : json['blockCode'],
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'blockStatus': !exists(json, 'blockStatus') ? undefined : json['blockStatus'],
        'blockName': !exists(json, 'blockName') ? undefined : json['blockName'],
    };
}

export function TourSeriesBlockTypeToJSON(value?: TourSeriesBlockType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'blockCode': value.blockCode,
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'blockStatus': value.blockStatus,
        'blockName': value.blockName,
    };
}

