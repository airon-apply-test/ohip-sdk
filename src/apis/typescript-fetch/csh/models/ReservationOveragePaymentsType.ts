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
import type { ReservationOveragePaymentType } from './ReservationOveragePaymentType';
import {
    ReservationOveragePaymentTypeFromJSON,
    ReservationOveragePaymentTypeFromJSONTyped,
    ReservationOveragePaymentTypeToJSON,
} from './ReservationOveragePaymentType';

/**
 * List of Reservation details for payment that has a folio window balance equal or higher to the credit limit set for the credit card payment method of that folio window.
 * @export
 * @interface ReservationOveragePaymentsType
 */
export interface ReservationOveragePaymentsType {
    /**
     * Reservation details to initiate the Credit Limit Overage process
     * @type {Array<ReservationOveragePaymentType>}
     * @memberof ReservationOveragePaymentsType
     */
    reservationOveragePayment?: Array<ReservationOveragePaymentType>;
    /**
     * Identifies the hotel code.
     * @type {string}
     * @memberof ReservationOveragePaymentsType
     */
    hotelId?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof ReservationOveragePaymentsType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the ReservationOveragePaymentsType interface.
 */
export function instanceOfReservationOveragePaymentsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationOveragePaymentsTypeFromJSON(json: any): ReservationOveragePaymentsType {
    return ReservationOveragePaymentsTypeFromJSONTyped(json, false);
}

export function ReservationOveragePaymentsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationOveragePaymentsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationOveragePayment': !exists(json, 'reservationOveragePayment') ? undefined : ((json['reservationOveragePayment'] as Array<any>).map(ReservationOveragePaymentTypeFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function ReservationOveragePaymentsTypeToJSON(value?: ReservationOveragePaymentsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationOveragePayment': value.reservationOveragePayment === undefined ? undefined : ((value.reservationOveragePayment as Array<any>).map(ReservationOveragePaymentTypeToJSON)),
        'hotelId': value.hotelId,
        'cashierId': value.cashierId,
    };
}

