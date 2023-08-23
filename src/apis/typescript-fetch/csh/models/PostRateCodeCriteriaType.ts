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
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { GuestCountsType } from './GuestCountsType';
import {
    GuestCountsTypeFromJSON,
    GuestCountsTypeFromJSONTyped,
    GuestCountsTypeToJSON,
} from './GuestCountsType';
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';

/**
 * Criteria type for posting a Rate Code amount to a guest folio.
 * @export
 * @interface PostRateCodeCriteriaType
 */
export interface PostRateCodeCriteriaType {
    /**
     * Hotel context for the Reservation.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    hotelId?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof PostRateCodeCriteriaType
     */
    reservationId?: ReservationId;
    /**
     * The Rate Code which is to be posted on the Guest Folio.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    rateCode?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof PostRateCodeCriteriaType
     */
    price?: CurrencyAmountType;
    /**
     * Posting quantity.
     * @type {number}
     * @memberof PostRateCodeCriteriaType
     */
    quantity?: number;
    /**
     * 
     * @type {GuestCountsType}
     * @memberof PostRateCodeCriteriaType
     */
    guestCounts?: GuestCountsType;
    /**
     * Posting remarks.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    remark?: string;
    /**
     * User-defined posting reference.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    reference?: string;
    /**
     * Check number for the posting.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    checkNo?: string;
    /**
     * Number of nights to post the Rate Code for. Used for Comp Accounting.
     * @type {number}
     * @memberof PostRateCodeCriteriaType
     */
    nights?: number;
    /**
     * Corrected arrangement code from the package associated to this transaction.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    arrangementCode?: string;
    /**
     * Adjustment code to be used for the manual posting of rate code.
     * @type {string}
     * @memberof PostRateCodeCriteriaType
     */
    adjustmentCode?: string;
    /**
     * Revenue Date or the business date of the posting.
     * @type {Date}
     * @memberof PostRateCodeCriteriaType
     */
    revenueDate?: Date;
    /**
     * The linked transaction number for this Posting of Rate Code.
     * @type {number}
     * @memberof PostRateCodeCriteriaType
     */
    parentTrxNo?: number;
    /**
     * Flag to indicate if Products(Packages) which are part of the Rate Code should be posted as part of this operation. If products are not required, the entire amount will be posted towards the room element of the Rate Code.
     * @type {boolean}
     * @memberof PostRateCodeCriteriaType
     */
    postProducts?: boolean;
    /**
     * 
     * @type {number}
     * @memberof PostRateCodeCriteriaType
     */
    folioWindowNo?: number;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof PostRateCodeCriteriaType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the PostRateCodeCriteriaType interface.
 */
export function instanceOfPostRateCodeCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostRateCodeCriteriaTypeFromJSON(json: any): PostRateCodeCriteriaType {
    return PostRateCodeCriteriaTypeFromJSONTyped(json, false);
}

export function PostRateCodeCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostRateCodeCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'rateCode': !exists(json, 'rateCode') ? undefined : json['rateCode'],
        'price': !exists(json, 'price') ? undefined : CurrencyAmountTypeFromJSON(json['price']),
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
        'guestCounts': !exists(json, 'guestCounts') ? undefined : GuestCountsTypeFromJSON(json['guestCounts']),
        'remark': !exists(json, 'remark') ? undefined : json['remark'],
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'checkNo': !exists(json, 'checkNo') ? undefined : json['checkNo'],
        'nights': !exists(json, 'nights') ? undefined : json['nights'],
        'arrangementCode': !exists(json, 'arrangementCode') ? undefined : json['arrangementCode'],
        'adjustmentCode': !exists(json, 'adjustmentCode') ? undefined : json['adjustmentCode'],
        'revenueDate': !exists(json, 'revenueDate') ? undefined : (new Date(json['revenueDate'])),
        'parentTrxNo': !exists(json, 'parentTrxNo') ? undefined : json['parentTrxNo'],
        'postProducts': !exists(json, 'postProducts') ? undefined : json['postProducts'],
        'folioWindowNo': !exists(json, 'folioWindowNo') ? undefined : json['folioWindowNo'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function PostRateCodeCriteriaTypeToJSON(value?: PostRateCodeCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'rateCode': value.rateCode,
        'price': CurrencyAmountTypeToJSON(value.price),
        'quantity': value.quantity,
        'guestCounts': GuestCountsTypeToJSON(value.guestCounts),
        'remark': value.remark,
        'reference': value.reference,
        'checkNo': value.checkNo,
        'nights': value.nights,
        'arrangementCode': value.arrangementCode,
        'adjustmentCode': value.adjustmentCode,
        'revenueDate': value.revenueDate === undefined ? undefined : (value.revenueDate.toISOString().substr(0,10)),
        'parentTrxNo': value.parentTrxNo,
        'postProducts': value.postProducts,
        'folioWindowNo': value.folioWindowNo,
        'cashierId': value.cashierId,
    };
}

