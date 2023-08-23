/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RoutingInstructionType } from './RoutingInstructionType';
import {
    RoutingInstructionTypeFromJSON,
    RoutingInstructionTypeFromJSONTyped,
    RoutingInstructionTypeToJSON,
} from './RoutingInstructionType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Room routing type.
 * @export
 * @interface RoutingInfoTypeRoom
 */
export interface RoutingInfoTypeRoom {
    /**
     * Room number to route the instructions.
     * @type {string}
     * @memberof RoutingInfoTypeRoom
     */
    roomId?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof RoutingInfoTypeRoom
     */
    guestNameId?: UniqueIDType;
    /**
     * Display Name for the guest.
     * @type {string}
     * @memberof RoutingInfoTypeRoom
     */
    guestDisplayName?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof RoutingInfoTypeRoom
     */
    reservationNameId?: UniqueIDType;
    /**
     * Set of routing instructions associated to this routing type.
     * @type {Array<RoutingInstructionType>}
     * @memberof RoutingInfoTypeRoom
     */
    instructions?: Array<RoutingInstructionType>;
}

/**
 * Check if a given object implements the RoutingInfoTypeRoom interface.
 */
export function instanceOfRoutingInfoTypeRoom(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoutingInfoTypeRoomFromJSON(json: any): RoutingInfoTypeRoom {
    return RoutingInfoTypeRoomFromJSONTyped(json, false);
}

export function RoutingInfoTypeRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInfoTypeRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomId': !exists(json, 'roomId') ? undefined : json['roomId'],
        'guestNameId': !exists(json, 'guestNameId') ? undefined : UniqueIDTypeFromJSON(json['guestNameId']),
        'guestDisplayName': !exists(json, 'guestDisplayName') ? undefined : json['guestDisplayName'],
        'reservationNameId': !exists(json, 'reservationNameId') ? undefined : UniqueIDTypeFromJSON(json['reservationNameId']),
        'instructions': !exists(json, 'instructions') ? undefined : ((json['instructions'] as Array<any>).map(RoutingInstructionTypeFromJSON)),
    };
}

export function RoutingInfoTypeRoomToJSON(value?: RoutingInfoTypeRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomId': value.roomId,
        'guestNameId': UniqueIDTypeToJSON(value.guestNameId),
        'guestDisplayName': value.guestDisplayName,
        'reservationNameId': UniqueIDTypeToJSON(value.reservationNameId),
        'instructions': value.instructions === undefined ? undefined : ((value.instructions as Array<any>).map(RoutingInstructionTypeToJSON)),
    };
}

