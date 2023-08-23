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
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';

/**
 * Criteria for transferring the Deposit Amount to the Guest Folio
 * @export
 * @interface PostDepositToGuestFolioType
 */
export interface PostDepositToGuestFolioType {
    /**
     * Resort for which the reservation defined.
     * @type {string}
     * @memberof PostDepositToGuestFolioType
     */
    hotelId?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof PostDepositToGuestFolioType
     */
    reservationId?: ReservationId;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof PostDepositToGuestFolioType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the PostDepositToGuestFolioType interface.
 */
export function instanceOfPostDepositToGuestFolioType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostDepositToGuestFolioTypeFromJSON(json: any): PostDepositToGuestFolioType {
    return PostDepositToGuestFolioTypeFromJSONTyped(json, false);
}

export function PostDepositToGuestFolioTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostDepositToGuestFolioType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function PostDepositToGuestFolioTypeToJSON(value?: PostDepositToGuestFolioType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'cashierId': value.cashierId,
    };
}

