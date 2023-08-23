/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { MasterRestrictionStatusesType } from './MasterRestrictionStatusesType';
import {
    MasterRestrictionStatusesTypeFromJSON,
    MasterRestrictionStatusesTypeFromJSONTyped,
    MasterRestrictionStatusesTypeToJSON,
} from './MasterRestrictionStatusesType';
import type { MembershipSearchType } from './MembershipSearchType';
import {
    MembershipSearchTypeFromJSON,
    MembershipSearchTypeFromJSONTyped,
    MembershipSearchTypeToJSON,
} from './MembershipSearchType';
import type { RoomStayType } from './RoomStayType';
import {
    RoomStayTypeFromJSON,
    RoomStayTypeFromJSONTyped,
    RoomStayTypeToJSON,
} from './RoomStayType';

/**
 * 
 * @export
 * @interface RateRoomDetailsType
 */
export interface RateRoomDetailsType {
    /**
     * Detail regarding customer loyalty program.
     * @type {Array<MembershipSearchType>}
     * @memberof RateRoomDetailsType
     */
    memberships?: Array<MembershipSearchType>;
    /**
     * 
     * @type {MasterRestrictionStatusesType}
     * @memberof RateRoomDetailsType
     */
    restrictionType?: MasterRestrictionStatusesType;
    /**
     * Room stay information.
     * @type {Array<RoomStayType>}
     * @memberof RateRoomDetailsType
     */
    roomStays?: Array<RoomStayType>;
}

/**
 * Check if a given object implements the RateRoomDetailsType interface.
 */
export function instanceOfRateRoomDetailsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RateRoomDetailsTypeFromJSON(json: any): RateRoomDetailsType {
    return RateRoomDetailsTypeFromJSONTyped(json, false);
}

export function RateRoomDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RateRoomDetailsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'memberships': !exists(json, 'memberships') ? undefined : ((json['memberships'] as Array<any>).map(MembershipSearchTypeFromJSON)),
        'restrictionType': !exists(json, 'restrictionType') ? undefined : MasterRestrictionStatusesTypeFromJSON(json['restrictionType']),
        'roomStays': !exists(json, 'roomStays') ? undefined : ((json['roomStays'] as Array<any>).map(RoomStayTypeFromJSON)),
    };
}

export function RateRoomDetailsTypeToJSON(value?: RateRoomDetailsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'memberships': value.memberships === undefined ? undefined : ((value.memberships as Array<any>).map(MembershipSearchTypeToJSON)),
        'restrictionType': MasterRestrictionStatusesTypeToJSON(value.restrictionType),
        'roomStays': value.roomStays === undefined ? undefined : ((value.roomStays as Array<any>).map(RoomStayTypeToJSON)),
    };
}

