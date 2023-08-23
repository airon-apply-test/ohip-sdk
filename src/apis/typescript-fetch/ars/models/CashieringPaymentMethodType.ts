/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AuthorizationInfoType } from './AuthorizationInfoType';
import {
    AuthorizationInfoTypeFromJSON,
    AuthorizationInfoTypeFromJSONTyped,
    AuthorizationInfoTypeToJSON,
} from './AuthorizationInfoType';
import type { AuthorizationRuleType } from './AuthorizationRuleType';
import {
    AuthorizationRuleTypeFromJSON,
    AuthorizationRuleTypeFromJSONTyped,
    AuthorizationRuleTypeToJSON,
} from './AuthorizationRuleType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { ResPaymentCardType } from './ResPaymentCardType';
import {
    ResPaymentCardTypeFromJSON,
    ResPaymentCardTypeFromJSONTyped,
    ResPaymentCardTypeToJSON,
} from './ResPaymentCardType';
import type { ReservationPaymentMethodTypeEmailFolioInfo } from './ReservationPaymentMethodTypeEmailFolioInfo';
import {
    ReservationPaymentMethodTypeEmailFolioInfoFromJSON,
    ReservationPaymentMethodTypeEmailFolioInfoFromJSONTyped,
    ReservationPaymentMethodTypeEmailFolioInfoToJSON,
} from './ReservationPaymentMethodTypeEmailFolioInfo';

/**
 * 
 * @export
 * @interface CashieringPaymentMethodType
 */
export interface CashieringPaymentMethodType {
    /**
     * 
     * @type {ResPaymentCardType}
     * @memberof CashieringPaymentMethodType
     */
    paymentCard?: ResPaymentCardType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CashieringPaymentMethodType
     */
    balance?: CurrencyAmountType;
    /**
     * 
     * @type {AuthorizationRuleType}
     * @memberof CashieringPaymentMethodType
     */
    authorizationRule?: AuthorizationRuleType;
    /**
     * 
     * @type {ReservationPaymentMethodTypeEmailFolioInfo}
     * @memberof CashieringPaymentMethodType
     */
    emailFolioInfo?: ReservationPaymentMethodTypeEmailFolioInfo;
    /**
     * 
     * @type {string}
     * @memberof CashieringPaymentMethodType
     */
    paymentMethod?: string;
    /**
     * 
     * @type {string}
     * @memberof CashieringPaymentMethodType
     */
    description?: string;
    /**
     * 
     * @type {number}
     * @memberof CashieringPaymentMethodType
     */
    folioView?: number;
    /**
     * 
     * @type {AuthorizationInfoType}
     * @memberof CashieringPaymentMethodType
     */
    authorizationApproval?: AuthorizationInfoType;
    /**
     * Track2 data for the card which would be sent to the CC vendor as part of the authorization.
     * @type {string}
     * @memberof CashieringPaymentMethodType
     */
    track2Data?: string;
}

/**
 * Check if a given object implements the CashieringPaymentMethodType interface.
 */
export function instanceOfCashieringPaymentMethodType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CashieringPaymentMethodTypeFromJSON(json: any): CashieringPaymentMethodType {
    return CashieringPaymentMethodTypeFromJSONTyped(json, false);
}

export function CashieringPaymentMethodTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashieringPaymentMethodType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'paymentCard': !exists(json, 'paymentCard') ? undefined : ResPaymentCardTypeFromJSON(json['paymentCard']),
        'balance': !exists(json, 'balance') ? undefined : CurrencyAmountTypeFromJSON(json['balance']),
        'authorizationRule': !exists(json, 'authorizationRule') ? undefined : AuthorizationRuleTypeFromJSON(json['authorizationRule']),
        'emailFolioInfo': !exists(json, 'emailFolioInfo') ? undefined : ReservationPaymentMethodTypeEmailFolioInfoFromJSON(json['emailFolioInfo']),
        'paymentMethod': !exists(json, 'paymentMethod') ? undefined : json['paymentMethod'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'folioView': !exists(json, 'folioView') ? undefined : json['folioView'],
        'authorizationApproval': !exists(json, 'authorizationApproval') ? undefined : AuthorizationInfoTypeFromJSON(json['authorizationApproval']),
        'track2Data': !exists(json, 'track2Data') ? undefined : json['track2Data'],
    };
}

export function CashieringPaymentMethodTypeToJSON(value?: CashieringPaymentMethodType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'paymentCard': ResPaymentCardTypeToJSON(value.paymentCard),
        'balance': CurrencyAmountTypeToJSON(value.balance),
        'authorizationRule': AuthorizationRuleTypeToJSON(value.authorizationRule),
        'emailFolioInfo': ReservationPaymentMethodTypeEmailFolioInfoToJSON(value.emailFolioInfo),
        'paymentMethod': value.paymentMethod,
        'description': value.description,
        'folioView': value.folioView,
        'authorizationApproval': AuthorizationInfoTypeToJSON(value.authorizationApproval),
        'track2Data': value.track2Data,
    };
}

