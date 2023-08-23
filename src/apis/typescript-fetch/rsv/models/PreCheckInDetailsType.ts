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
import type { ReservationArrivalInfoType } from './ReservationArrivalInfoType';
import {
    ReservationArrivalInfoTypeFromJSON,
    ReservationArrivalInfoTypeFromJSONTyped,
    ReservationArrivalInfoTypeToJSON,
} from './ReservationArrivalInfoType';
import type { ReservationPaymentMethodType } from './ReservationPaymentMethodType';
import {
    ReservationPaymentMethodTypeFromJSON,
    ReservationPaymentMethodTypeFromJSONTyped,
    ReservationPaymentMethodTypeToJSON,
} from './ReservationPaymentMethodType';
import type { TransportInfoType } from './TransportInfoType';
import {
    TransportInfoTypeFromJSON,
    TransportInfoTypeFromJSONTyped,
    TransportInfoTypeToJSON,
} from './TransportInfoType';

/**
 * Type describing the different details regarding a pre-registered reservation.
 * @export
 * @interface PreCheckInDetailsType
 */
export interface PreCheckInDetailsType {
    /**
     * 
     * @type {ReservationArrivalInfoType}
     * @memberof PreCheckInDetailsType
     */
    arrival?: ReservationArrivalInfoType;
    /**
     * Room Number to be assigned to the reservation.
     * @type {string}
     * @memberof PreCheckInDetailsType
     */
    roomId?: string;
    /**
     * Currency code preferred by guest.
     * @type {string}
     * @memberof PreCheckInDetailsType
     */
    guestPreferredCurrency?: string;
    /**
     * Attribute AllowMobileViewFolio is set to true when the reservation is eligible for viewing folio using mobile device.
     * @type {boolean}
     * @memberof PreCheckInDetailsType
     */
    allowMobileViewFolio?: boolean;
    /**
     * Defines reservation payment methods.
     * @type {Array<ReservationPaymentMethodType>}
     * @memberof PreCheckInDetailsType
     */
    reservationPaymentMethods?: Array<ReservationPaymentMethodType>;
    /**
     * 
     * @type {TransportInfoType}
     * @memberof PreCheckInDetailsType
     */
    arrivalTransport?: TransportInfoType;
}

/**
 * Check if a given object implements the PreCheckInDetailsType interface.
 */
export function instanceOfPreCheckInDetailsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PreCheckInDetailsTypeFromJSON(json: any): PreCheckInDetailsType {
    return PreCheckInDetailsTypeFromJSONTyped(json, false);
}

export function PreCheckInDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PreCheckInDetailsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'arrival': !exists(json, 'arrival') ? undefined : ReservationArrivalInfoTypeFromJSON(json['arrival']),
        'roomId': !exists(json, 'roomId') ? undefined : json['roomId'],
        'guestPreferredCurrency': !exists(json, 'guestPreferredCurrency') ? undefined : json['guestPreferredCurrency'],
        'allowMobileViewFolio': !exists(json, 'allowMobileViewFolio') ? undefined : json['allowMobileViewFolio'],
        'reservationPaymentMethods': !exists(json, 'reservationPaymentMethods') ? undefined : ((json['reservationPaymentMethods'] as Array<any>).map(ReservationPaymentMethodTypeFromJSON)),
        'arrivalTransport': !exists(json, 'arrivalTransport') ? undefined : TransportInfoTypeFromJSON(json['arrivalTransport']),
    };
}

export function PreCheckInDetailsTypeToJSON(value?: PreCheckInDetailsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'arrival': ReservationArrivalInfoTypeToJSON(value.arrival),
        'roomId': value.roomId,
        'guestPreferredCurrency': value.guestPreferredCurrency,
        'allowMobileViewFolio': value.allowMobileViewFolio,
        'reservationPaymentMethods': value.reservationPaymentMethods === undefined ? undefined : ((value.reservationPaymentMethods as Array<any>).map(ReservationPaymentMethodTypeToJSON)),
        'arrivalTransport': TransportInfoTypeToJSON(value.arrivalTransport),
    };
}

