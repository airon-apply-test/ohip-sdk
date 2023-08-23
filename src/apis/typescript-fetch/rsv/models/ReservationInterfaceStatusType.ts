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
import type { HotelInterfaceType } from './HotelInterfaceType';
import {
    HotelInterfaceTypeFromJSON,
    HotelInterfaceTypeFromJSONTyped,
    HotelInterfaceTypeToJSON,
} from './HotelInterfaceType';
import type { InterfaceRightsStatusType } from './InterfaceRightsStatusType';
import {
    InterfaceRightsStatusTypeFromJSON,
    InterfaceRightsStatusTypeFromJSONTyped,
    InterfaceRightsStatusTypeToJSON,
} from './InterfaceRightsStatusType';

/**
 * Hotel Interface Type for a reservation and status of the various services
 * @export
 * @interface ReservationInterfaceStatusType
 */
export interface ReservationInterfaceStatusType {
    /**
     * Identifier for the room extension
     * @type {string}
     * @memberof ReservationInterfaceStatusType
     */
    roomExtension?: string;
    /**
     * 
     * @type {HotelInterfaceType}
     * @memberof ReservationInterfaceStatusType
     */
    hotelInterface?: HotelInterfaceType;
    /**
     * Contains a list of status/rights for the various services under this interface
     * @type {Array<InterfaceRightsStatusType>}
     * @memberof ReservationInterfaceStatusType
     */
    interfaceRights?: Array<InterfaceRightsStatusType>;
}

/**
 * Check if a given object implements the ReservationInterfaceStatusType interface.
 */
export function instanceOfReservationInterfaceStatusType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationInterfaceStatusTypeFromJSON(json: any): ReservationInterfaceStatusType {
    return ReservationInterfaceStatusTypeFromJSONTyped(json, false);
}

export function ReservationInterfaceStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationInterfaceStatusType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomExtension': !exists(json, 'roomExtension') ? undefined : json['roomExtension'],
        'hotelInterface': !exists(json, 'hotelInterface') ? undefined : HotelInterfaceTypeFromJSON(json['hotelInterface']),
        'interfaceRights': !exists(json, 'interfaceRights') ? undefined : ((json['interfaceRights'] as Array<any>).map(InterfaceRightsStatusTypeFromJSON)),
    };
}

export function ReservationInterfaceStatusTypeToJSON(value?: ReservationInterfaceStatusType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomExtension': value.roomExtension,
        'hotelInterface': HotelInterfaceTypeToJSON(value.hotelInterface),
        'interfaceRights': value.interfaceRights === undefined ? undefined : ((value.interfaceRights as Array<any>).map(InterfaceRightsStatusTypeToJSON)),
    };
}

