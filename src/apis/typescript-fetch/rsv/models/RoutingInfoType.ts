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
import type { RoutingInfoTypeComp } from './RoutingInfoTypeComp';
import {
    RoutingInfoTypeCompFromJSON,
    RoutingInfoTypeCompFromJSONTyped,
    RoutingInfoTypeCompToJSON,
} from './RoutingInfoTypeComp';
import type { RoutingInfoTypeFolio } from './RoutingInfoTypeFolio';
import {
    RoutingInfoTypeFolioFromJSON,
    RoutingInfoTypeFolioFromJSONTyped,
    RoutingInfoTypeFolioToJSON,
} from './RoutingInfoTypeFolio';
import type { RoutingInfoTypeRequest } from './RoutingInfoTypeRequest';
import {
    RoutingInfoTypeRequestFromJSON,
    RoutingInfoTypeRequestFromJSONTyped,
    RoutingInfoTypeRequestToJSON,
} from './RoutingInfoTypeRequest';
import type { RoutingInfoTypeRoom } from './RoutingInfoTypeRoom';
import {
    RoutingInfoTypeRoomFromJSON,
    RoutingInfoTypeRoomFromJSONTyped,
    RoutingInfoTypeRoomToJSON,
} from './RoutingInfoTypeRoom';

/**
 * A routing info object can either be of type Folio OR of type Room with its corresponding object.
 * @export
 * @interface RoutingInfoType
 */
export interface RoutingInfoType {
    /**
     * 
     * @type {RoutingInfoTypeFolio}
     * @memberof RoutingInfoType
     */
    folio?: RoutingInfoTypeFolio;
    /**
     * 
     * @type {RoutingInfoTypeRoom}
     * @memberof RoutingInfoType
     */
    room?: RoutingInfoTypeRoom;
    /**
     * 
     * @type {RoutingInfoTypeComp}
     * @memberof RoutingInfoType
     */
    comp?: RoutingInfoTypeComp;
    /**
     * 
     * @type {RoutingInfoTypeRequest}
     * @memberof RoutingInfoType
     */
    request?: RoutingInfoTypeRequest;
    /**
     * On a successful update, the transactions that are already posted in the guest's folio will be re-organized based on the configured instructions.
     * @type {boolean}
     * @memberof RoutingInfoType
     */
    refreshFolio?: boolean;
}

/**
 * Check if a given object implements the RoutingInfoType interface.
 */
export function instanceOfRoutingInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoutingInfoTypeFromJSON(json: any): RoutingInfoType {
    return RoutingInfoTypeFromJSONTyped(json, false);
}

export function RoutingInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'folio': !exists(json, 'folio') ? undefined : RoutingInfoTypeFolioFromJSON(json['folio']),
        'room': !exists(json, 'room') ? undefined : RoutingInfoTypeRoomFromJSON(json['room']),
        'comp': !exists(json, 'comp') ? undefined : RoutingInfoTypeCompFromJSON(json['comp']),
        'request': !exists(json, 'request') ? undefined : RoutingInfoTypeRequestFromJSON(json['request']),
        'refreshFolio': !exists(json, 'refreshFolio') ? undefined : json['refreshFolio'],
    };
}

export function RoutingInfoTypeToJSON(value?: RoutingInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'folio': RoutingInfoTypeFolioToJSON(value.folio),
        'room': RoutingInfoTypeRoomToJSON(value.room),
        'comp': RoutingInfoTypeCompToJSON(value.comp),
        'request': RoutingInfoTypeRequestToJSON(value.request),
        'refreshFolio': value.refreshFolio,
    };
}

