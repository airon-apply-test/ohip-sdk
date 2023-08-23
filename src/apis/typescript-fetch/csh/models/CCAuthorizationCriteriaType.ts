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
import type { CardAuthorizationTransactionType } from './CardAuthorizationTransactionType';
import {
    CardAuthorizationTransactionTypeFromJSON,
    CardAuthorizationTransactionTypeFromJSONTyped,
    CardAuthorizationTransactionTypeToJSON,
} from './CardAuthorizationTransactionType';
import type { CashieringPaymentMethodType } from './CashieringPaymentMethodType';
import {
    CashieringPaymentMethodTypeFromJSON,
    CashieringPaymentMethodTypeFromJSONTyped,
    CashieringPaymentMethodTypeToJSON,
} from './CashieringPaymentMethodType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { ProfileId } from './ProfileId';
import {
    ProfileIdFromJSON,
    ProfileIdFromJSONTyped,
    ProfileIdToJSON,
} from './ProfileId';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Contains the credit card criteria for authorization.
 * @export
 * @interface CCAuthorizationCriteriaType
 */
export interface CCAuthorizationCriteriaType {
    /**
     * Identifies the hotel code to authorize a credit card amount for.
     * @type {string}
     * @memberof CCAuthorizationCriteriaType
     */
    hotelId?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof CCAuthorizationCriteriaType
     */
    reservationId?: UniqueIDType;
    /**
     * 
     * @type {ProfileId}
     * @memberof CCAuthorizationCriteriaType
     */
    profileId?: ProfileId;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CCAuthorizationCriteriaType
     */
    amount?: CurrencyAmountType;
    /**
     * The folio window where this authorization is applied.
     * @type {number}
     * @memberof CCAuthorizationCriteriaType
     */
    folioView?: number;
    /**
     * 
     * @type {CashieringPaymentMethodType}
     * @memberof CCAuthorizationCriteriaType
     */
    payment?: CashieringPaymentMethodType;
    /**
     * Applicable for chip and pin. The ID of the terminal where the chip and pin device is connected.
     * @type {string}
     * @memberof CCAuthorizationCriteriaType
     */
    terminalId?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CCAuthorizationCriteriaType
     */
    incidentalAmount?: CurrencyAmountType;
    /**
     * Flag to indicate if the approval amount should be calculated before authorization. If this is true and if the amount is 0 the approval amount will be calculated based on the rules.
     * @type {boolean}
     * @memberof CCAuthorizationCriteriaType
     */
    evaluateApprovalAmount?: boolean;
    /**
     * 
     * @type {CardAuthorizationTransactionType}
     * @memberof CCAuthorizationCriteriaType
     */
    sourceOfAuthorization?: CardAuthorizationTransactionType;
    /**
     * Update the card details on the reservations for Chip and Pin Authorizations.
     * @type {boolean}
     * @memberof CCAuthorizationCriteriaType
     */
    updateReservation?: boolean;
    /**
     * Session Id registered in the WebSocket component for asynchronous Credit Card handling.
     * @type {string}
     * @memberof CCAuthorizationCriteriaType
     */
    ccRequestId?: string;
}

/**
 * Check if a given object implements the CCAuthorizationCriteriaType interface.
 */
export function instanceOfCCAuthorizationCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CCAuthorizationCriteriaTypeFromJSON(json: any): CCAuthorizationCriteriaType {
    return CCAuthorizationCriteriaTypeFromJSONTyped(json, false);
}

export function CCAuthorizationCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CCAuthorizationCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : UniqueIDTypeFromJSON(json['reservationId']),
        'profileId': !exists(json, 'profileId') ? undefined : ProfileIdFromJSON(json['profileId']),
        'amount': !exists(json, 'amount') ? undefined : CurrencyAmountTypeFromJSON(json['amount']),
        'folioView': !exists(json, 'folioView') ? undefined : json['folioView'],
        'payment': !exists(json, 'payment') ? undefined : CashieringPaymentMethodTypeFromJSON(json['payment']),
        'terminalId': !exists(json, 'terminalId') ? undefined : json['terminalId'],
        'incidentalAmount': !exists(json, 'incidentalAmount') ? undefined : CurrencyAmountTypeFromJSON(json['incidentalAmount']),
        'evaluateApprovalAmount': !exists(json, 'evaluateApprovalAmount') ? undefined : json['evaluateApprovalAmount'],
        'sourceOfAuthorization': !exists(json, 'sourceOfAuthorization') ? undefined : CardAuthorizationTransactionTypeFromJSON(json['sourceOfAuthorization']),
        'updateReservation': !exists(json, 'updateReservation') ? undefined : json['updateReservation'],
        'ccRequestId': !exists(json, 'ccRequestId') ? undefined : json['ccRequestId'],
    };
}

export function CCAuthorizationCriteriaTypeToJSON(value?: CCAuthorizationCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationId': UniqueIDTypeToJSON(value.reservationId),
        'profileId': ProfileIdToJSON(value.profileId),
        'amount': CurrencyAmountTypeToJSON(value.amount),
        'folioView': value.folioView,
        'payment': CashieringPaymentMethodTypeToJSON(value.payment),
        'terminalId': value.terminalId,
        'incidentalAmount': CurrencyAmountTypeToJSON(value.incidentalAmount),
        'evaluateApprovalAmount': value.evaluateApprovalAmount,
        'sourceOfAuthorization': CardAuthorizationTransactionTypeToJSON(value.sourceOfAuthorization),
        'updateReservation': value.updateReservation,
        'ccRequestId': value.ccRequestId,
    };
}

