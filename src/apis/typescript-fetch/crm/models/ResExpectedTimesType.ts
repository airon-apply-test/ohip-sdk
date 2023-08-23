/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Holds the Arrival and Departure Time Information
 * @export
 * @interface ResExpectedTimesType
 */
export interface ResExpectedTimesType {
    /**
     * Arrival Time
     * @type {Date}
     * @memberof ResExpectedTimesType
     */
    reservationExpectedArrivalTime?: Date;
    /**
     * Departure Time
     * @type {Date}
     * @memberof ResExpectedTimesType
     */
    reservationExpectedDepartureTime?: Date;
}

/**
 * Check if a given object implements the ResExpectedTimesType interface.
 */
export function instanceOfResExpectedTimesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResExpectedTimesTypeFromJSON(json: any): ResExpectedTimesType {
    return ResExpectedTimesTypeFromJSONTyped(json, false);
}

export function ResExpectedTimesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResExpectedTimesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationExpectedArrivalTime': !exists(json, 'reservationExpectedArrivalTime') ? undefined : (new Date(json['reservationExpectedArrivalTime'])),
        'reservationExpectedDepartureTime': !exists(json, 'reservationExpectedDepartureTime') ? undefined : (new Date(json['reservationExpectedDepartureTime'])),
    };
}

export function ResExpectedTimesTypeToJSON(value?: ResExpectedTimesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationExpectedArrivalTime': value.reservationExpectedArrivalTime === undefined ? undefined : (value.reservationExpectedArrivalTime.toISOString()),
        'reservationExpectedDepartureTime': value.reservationExpectedDepartureTime === undefined ? undefined : (value.reservationExpectedDepartureTime.toISOString()),
    };
}

