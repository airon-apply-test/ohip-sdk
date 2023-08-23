/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CardNumberTypeType } from './CardNumberTypeType';
import {
    CardNumberTypeTypeFromJSON,
    CardNumberTypeTypeFromJSONTyped,
    CardNumberTypeTypeToJSON,
} from './CardNumberTypeType';
import type { CardProcessingType } from './CardProcessingType';
import {
    CardProcessingTypeFromJSON,
    CardProcessingTypeFromJSONTyped,
    CardProcessingTypeToJSON,
} from './CardProcessingType';
import type { CardTypeType } from './CardTypeType';
import {
    CardTypeTypeFromJSON,
    CardTypeTypeFromJSONTyped,
    CardTypeTypeToJSON,
} from './CardTypeType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Information on a credit card for the customer.
 * @export
 * @interface ResPaymentCardType
 */
export interface ResPaymentCardType {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ResPaymentCardType
     */
    cardId?: UniqueIDType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof ResPaymentCardType
     */
    currentAuthorizedAmount?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof ResPaymentCardType
     */
    approvalAmountNeeded?: CurrencyAmountType;
    /**
     * 
     * @type {CardTypeType}
     * @memberof ResPaymentCardType
     */
    cardType?: CardTypeType;
    /**
     * Indicates the user defined credit card type if credit card type from a defined list is not provided
     * @type {string}
     * @memberof ResPaymentCardType
     */
    userDefinedCardType?: string;
    /**
     * 
     * @type {string}
     * @memberof ResPaymentCardType
     */
    cardNumber?: string;
    /**
     * 
     * @type {string}
     * @memberof ResPaymentCardType
     */
    cardNumberMasked?: string;
    /**
     * 
     * @type {string}
     * @memberof ResPaymentCardType
     */
    cardNumberLast4Digits?: string;
    /**
     * Expiration date of the credit card
     * @type {Date}
     * @memberof ResPaymentCardType
     */
    expirationDate?: Date;
    /**
     * Masked Expiration date of the credit card
     * @type {string}
     * @memberof ResPaymentCardType
     */
    expirationDateMasked?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ResPaymentCardType
     */
    expirationDateExpired?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ResPaymentCardType
     */
    cardHolderName?: string;
    /**
     * Attach the credit card to profile.
     * @type {boolean}
     * @memberof ResPaymentCardType
     */
    attachCreditCardToProfile?: boolean;
    /**
     * 
     * @type {CardProcessingType}
     * @memberof ResPaymentCardType
     */
    processing?: CardProcessingType;
    /**
     * 
     * @type {boolean}
     * @memberof ResPaymentCardType
     */
    swiped?: boolean;
    /**
     * 
     * @type {CardNumberTypeType}
     * @memberof ResPaymentCardType
     */
    cardNumberType?: CardNumberTypeType;
}

/**
 * Check if a given object implements the ResPaymentCardType interface.
 */
export function instanceOfResPaymentCardType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResPaymentCardTypeFromJSON(json: any): ResPaymentCardType {
    return ResPaymentCardTypeFromJSONTyped(json, false);
}

export function ResPaymentCardTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResPaymentCardType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cardId': !exists(json, 'cardId') ? undefined : UniqueIDTypeFromJSON(json['cardId']),
        'currentAuthorizedAmount': !exists(json, 'currentAuthorizedAmount') ? undefined : CurrencyAmountTypeFromJSON(json['currentAuthorizedAmount']),
        'approvalAmountNeeded': !exists(json, 'approvalAmountNeeded') ? undefined : CurrencyAmountTypeFromJSON(json['approvalAmountNeeded']),
        'cardType': !exists(json, 'cardType') ? undefined : CardTypeTypeFromJSON(json['cardType']),
        'userDefinedCardType': !exists(json, 'userDefinedCardType') ? undefined : json['userDefinedCardType'],
        'cardNumber': !exists(json, 'cardNumber') ? undefined : json['cardNumber'],
        'cardNumberMasked': !exists(json, 'cardNumberMasked') ? undefined : json['cardNumberMasked'],
        'cardNumberLast4Digits': !exists(json, 'cardNumberLast4Digits') ? undefined : json['cardNumberLast4Digits'],
        'expirationDate': !exists(json, 'expirationDate') ? undefined : (new Date(json['expirationDate'])),
        'expirationDateMasked': !exists(json, 'expirationDateMasked') ? undefined : json['expirationDateMasked'],
        'expirationDateExpired': !exists(json, 'expirationDateExpired') ? undefined : json['expirationDateExpired'],
        'cardHolderName': !exists(json, 'cardHolderName') ? undefined : json['cardHolderName'],
        'attachCreditCardToProfile': !exists(json, 'attachCreditCardToProfile') ? undefined : json['attachCreditCardToProfile'],
        'processing': !exists(json, 'processing') ? undefined : CardProcessingTypeFromJSON(json['processing']),
        'swiped': !exists(json, 'swiped') ? undefined : json['swiped'],
        'cardNumberType': !exists(json, 'cardNumberType') ? undefined : CardNumberTypeTypeFromJSON(json['cardNumberType']),
    };
}

export function ResPaymentCardTypeToJSON(value?: ResPaymentCardType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cardId': UniqueIDTypeToJSON(value.cardId),
        'currentAuthorizedAmount': CurrencyAmountTypeToJSON(value.currentAuthorizedAmount),
        'approvalAmountNeeded': CurrencyAmountTypeToJSON(value.approvalAmountNeeded),
        'cardType': CardTypeTypeToJSON(value.cardType),
        'userDefinedCardType': value.userDefinedCardType,
        'cardNumber': value.cardNumber,
        'cardNumberMasked': value.cardNumberMasked,
        'cardNumberLast4Digits': value.cardNumberLast4Digits,
        'expirationDate': value.expirationDate === undefined ? undefined : (value.expirationDate.toISOString().substr(0,10)),
        'expirationDateMasked': value.expirationDateMasked,
        'expirationDateExpired': value.expirationDateExpired,
        'cardHolderName': value.cardHolderName,
        'attachCreditCardToProfile': value.attachCreditCardToProfile,
        'processing': CardProcessingTypeToJSON(value.processing),
        'swiped': value.swiped,
        'cardNumberType': CardNumberTypeTypeToJSON(value.cardNumberType),
    };
}

