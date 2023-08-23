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
import type { PrepaidCardRedemptionType } from './PrepaidCardRedemptionType';
import {
    PrepaidCardRedemptionTypeFromJSON,
    PrepaidCardRedemptionTypeFromJSONTyped,
    PrepaidCardRedemptionTypeToJSON,
} from './PrepaidCardRedemptionType';
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';

/**
 * Criteria for Prepaid card redeem request.
 * @export
 * @interface RedeemPrepaidCardCriteriaType
 */
export interface RedeemPrepaidCardCriteriaType {
    /**
     * Hotel code.
     * @type {string}
     * @memberof RedeemPrepaidCardCriteriaType
     */
    hotelId?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof RedeemPrepaidCardCriteriaType
     */
    reservationId?: ReservationId;
    /**
     * A prepaid redemption info object to be used for posting a payment.
     * @type {Array<PrepaidCardRedemptionType>}
     * @memberof RedeemPrepaidCardCriteriaType
     */
    prepaidCardRedemptions?: Array<PrepaidCardRedemptionType>;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof RedeemPrepaidCardCriteriaType
     */
    cashierId?: number;
    /**
     * Folio Window.
     * @type {number}
     * @memberof RedeemPrepaidCardCriteriaType
     */
    folioView?: number;
}

/**
 * Check if a given object implements the RedeemPrepaidCardCriteriaType interface.
 */
export function instanceOfRedeemPrepaidCardCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RedeemPrepaidCardCriteriaTypeFromJSON(json: any): RedeemPrepaidCardCriteriaType {
    return RedeemPrepaidCardCriteriaTypeFromJSONTyped(json, false);
}

export function RedeemPrepaidCardCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RedeemPrepaidCardCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'prepaidCardRedemptions': !exists(json, 'prepaidCardRedemptions') ? undefined : ((json['prepaidCardRedemptions'] as Array<any>).map(PrepaidCardRedemptionTypeFromJSON)),
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
        'folioView': !exists(json, 'folioView') ? undefined : json['folioView'],
    };
}

export function RedeemPrepaidCardCriteriaTypeToJSON(value?: RedeemPrepaidCardCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'prepaidCardRedemptions': value.prepaidCardRedemptions === undefined ? undefined : ((value.prepaidCardRedemptions as Array<any>).map(PrepaidCardRedemptionTypeToJSON)),
        'cashierId': value.cashierId,
        'folioView': value.folioView,
    };
}

