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
 * Allows for a choice in description of the amount of time spanned by this type. EndDate specifies a specific date, while Duration provides a measure of time to add to the StartDate to yield end date.
 * @export
 * @interface DateTimeSpanType
 */
export interface DateTimeSpanType {
    /**
     * 
     * @type {Date}
     * @memberof DateTimeSpanType
     */
    startDateTime?: Date;
    /**
     * 
     * @type {Date}
     * @memberof DateTimeSpanType
     */
    endDateTime?: Date;
}

/**
 * Check if a given object implements the DateTimeSpanType interface.
 */
export function instanceOfDateTimeSpanType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DateTimeSpanTypeFromJSON(json: any): DateTimeSpanType {
    return DateTimeSpanTypeFromJSONTyped(json, false);
}

export function DateTimeSpanTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DateTimeSpanType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'startDateTime': !exists(json, 'startDateTime') ? undefined : (new Date(json['startDateTime'])),
        'endDateTime': !exists(json, 'endDateTime') ? undefined : (new Date(json['endDateTime'])),
    };
}

export function DateTimeSpanTypeToJSON(value?: DateTimeSpanType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'startDateTime': value.startDateTime === undefined ? undefined : (value.startDateTime.toISOString()),
        'endDateTime': value.endDateTime === undefined ? undefined : (value.endDateTime.toISOString()),
    };
}

