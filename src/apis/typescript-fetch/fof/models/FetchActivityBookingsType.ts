/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ActivityListInner } from './ActivityListInner';
import {
    ActivityListInnerFromJSON,
    ActivityListInnerFromJSONTyped,
    ActivityListInnerToJSON,
} from './ActivityListInner';
import type { AddressType } from './AddressType';
import {
    AddressTypeFromJSON,
    AddressTypeFromJSONTyped,
    AddressTypeToJSON,
} from './AddressType';
import type { PersonNameType } from './PersonNameType';
import {
    PersonNameTypeFromJSON,
    PersonNameTypeFromJSONTyped,
    PersonNameTypeToJSON,
} from './PersonNameType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Used in the request message to describe the "filtering Criteria" when executing an activity lookup.
 * @export
 * @interface FetchActivityBookingsType
 */
export interface FetchActivityBookingsType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof FetchActivityBookingsType
     */
    profileId?: Array<UniqueIDType>;
    /**
     * 
     * @type {PersonNameType}
     * @memberof FetchActivityBookingsType
     */
    personName?: PersonNameType;
    /**
     * 
     * @type {AddressType}
     * @memberof FetchActivityBookingsType
     */
    address?: AddressType;
    /**
     * A collection of Activity objects.
     * @type {Array<ActivityListInner>}
     * @memberof FetchActivityBookingsType
     */
    activities?: Array<ActivityListInner>;
    /**
     * Hotel Code, It is used to filter hotel specific children to this specific hotel code.
     * @type {string}
     * @memberof FetchActivityBookingsType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the FetchActivityBookingsType interface.
 */
export function instanceOfFetchActivityBookingsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FetchActivityBookingsTypeFromJSON(json: any): FetchActivityBookingsType {
    return FetchActivityBookingsTypeFromJSONTyped(json, false);
}

export function FetchActivityBookingsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FetchActivityBookingsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileId': !exists(json, 'profileId') ? undefined : ((json['profileId'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'personName': !exists(json, 'personName') ? undefined : PersonNameTypeFromJSON(json['personName']),
        'address': !exists(json, 'address') ? undefined : AddressTypeFromJSON(json['address']),
        'activities': !exists(json, 'activities') ? undefined : ((json['activities'] as Array<any>).map(ActivityListInnerFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function FetchActivityBookingsTypeToJSON(value?: FetchActivityBookingsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileId': value.profileId === undefined ? undefined : ((value.profileId as Array<any>).map(UniqueIDTypeToJSON)),
        'personName': PersonNameTypeToJSON(value.personName),
        'address': AddressTypeToJSON(value.address),
        'activities': value.activities === undefined ? undefined : ((value.activities as Array<any>).map(ActivityListInnerToJSON)),
        'hotelId': value.hotelId,
    };
}

