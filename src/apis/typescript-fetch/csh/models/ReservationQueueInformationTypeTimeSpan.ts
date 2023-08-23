/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Allows for a choice in description of the amount of time spanned by this type. EndDate specifies a specific date, while Duration provides a measure of time to add to the StartDate to yield end date.
 * @export
 * @interface ReservationQueueInformationTypeTimeSpan
 */
export interface ReservationQueueInformationTypeTimeSpan {
    /**
     * 
     * @type {Date}
     * @memberof ReservationQueueInformationTypeTimeSpan
     */
    startDate?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ReservationQueueInformationTypeTimeSpan
     */
    endDate?: Date;
    /**
     * 
     * @type {string}
     * @memberof ReservationQueueInformationTypeTimeSpan
     */
    duration?: string;
    /**
     * The total duration, in seconds, the reservation is on Queue.
     * @type {number}
     * @memberof ReservationQueueInformationTypeTimeSpan
     */
    durationInSeconds?: number;
}

/**
 * Check if a given object implements the ReservationQueueInformationTypeTimeSpan interface.
 */
export function instanceOfReservationQueueInformationTypeTimeSpan(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationQueueInformationTypeTimeSpanFromJSON(json: any): ReservationQueueInformationTypeTimeSpan {
    return ReservationQueueInformationTypeTimeSpanFromJSONTyped(json, false);
}

export function ReservationQueueInformationTypeTimeSpanFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationQueueInformationTypeTimeSpan {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !exists(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'duration': !exists(json, 'duration') ? undefined : json['duration'],
        'durationInSeconds': !exists(json, 'durationInSeconds') ? undefined : json['durationInSeconds'],
    };
}

export function ReservationQueueInformationTypeTimeSpanToJSON(value?: ReservationQueueInformationTypeTimeSpan | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0,10)),
        'duration': value.duration,
        'durationInSeconds': value.durationInSeconds,
    };
}

