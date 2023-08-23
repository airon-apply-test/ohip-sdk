/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { FBAStatusType } from './FBAStatusType';
import {
    FBAStatusTypeFromJSON,
    FBAStatusTypeFromJSONTyped,
    FBAStatusTypeToJSON,
} from './FBAStatusType';

/**
 * FBA ( Flexible Benefits Awards ) related fields.
 * @export
 * @interface FBAInfoType
 */
export interface FBAInfoType {
    /**
     * 
     * @type {FBAStatusType}
     * @memberof FBAInfoType
     */
    status?: FBAStatusType;
    /**
     * Award's FBA monetary values.
     * @type {number}
     * @memberof FBAInfoType
     */
    monetaryValue?: number;
    /**
     * Award's FBA amount.
     * @type {number}
     * @memberof FBAInfoType
     */
    amount?: number;
    /**
     * Award's FBA posted amount.
     * @type {number}
     * @memberof FBAInfoType
     */
    postedAmount?: number;
    /**
     * Award's FBA reimbursed amount.
     * @type {number}
     * @memberof FBAInfoType
     */
    reimbursedAmount?: number;
    /**
     * Date and time of the FBA posting.
     * @type {Date}
     * @memberof FBAInfoType
     */
    postingDateTime?: Date;
    /**
     * Business date of the FBA posting.
     * @type {Date}
     * @memberof FBAInfoType
     */
    postingBusinessDate?: Date;
    /**
     * Date and time of the FBA settlement.
     * @type {Date}
     * @memberof FBAInfoType
     */
    settlementDateTime?: Date;
    /**
     * Business date of the FBA settlement.
     * @type {Date}
     * @memberof FBAInfoType
     */
    settlementBusinessDate?: Date;
    /**
     * Date and time of the FBA reimbursement.
     * @type {Date}
     * @memberof FBAInfoType
     */
    reimbursementDateTime?: Date;
    /**
     * Business date of the FBA reimbursement.
     * @type {Date}
     * @memberof FBAInfoType
     */
    reimbursementBusinessDate?: Date;
    /**
     * Business date of the FBA bill generation.
     * @type {Date}
     * @memberof FBAInfoType
     */
    fbaBillGenDate?: Date;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof FBAInfoType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof FBAInfoType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof FBAInfoType
     */
    decimalPlaces?: number;
    /**
     * Indicates if this certificate is a Flexible Benefit Award certificate.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    flexibleBenefitAward?: boolean;
    /**
     * Indicates whether FBA has been posted.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    posted?: boolean;
    /**
     * Indicates whether FBA has been settled.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    settled?: boolean;
    /**
     * Indicates whether FBA has been reimbursed.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    reimbursed?: boolean;
    /**
     * Marks if the certificate is eligible for resettlement
     * @type {boolean}
     * @memberof FBAInfoType
     */
    resettleAllowed?: boolean;
    /**
     * Marks if the certificate is eligible for reimbursement.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    reimburseAllowed?: boolean;
    /**
     * Indicates whether the certificate is Orphan or not.
     * @type {boolean}
     * @memberof FBAInfoType
     */
    orphanCertificate?: boolean;
}

/**
 * Check if a given object implements the FBAInfoType interface.
 */
export function instanceOfFBAInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FBAInfoTypeFromJSON(json: any): FBAInfoType {
    return FBAInfoTypeFromJSONTyped(json, false);
}

export function FBAInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FBAInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'status': !exists(json, 'status') ? undefined : FBAStatusTypeFromJSON(json['status']),
        'monetaryValue': !exists(json, 'monetaryValue') ? undefined : json['monetaryValue'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'postedAmount': !exists(json, 'postedAmount') ? undefined : json['postedAmount'],
        'reimbursedAmount': !exists(json, 'reimbursedAmount') ? undefined : json['reimbursedAmount'],
        'postingDateTime': !exists(json, 'postingDateTime') ? undefined : (new Date(json['postingDateTime'])),
        'postingBusinessDate': !exists(json, 'postingBusinessDate') ? undefined : (new Date(json['postingBusinessDate'])),
        'settlementDateTime': !exists(json, 'settlementDateTime') ? undefined : (new Date(json['settlementDateTime'])),
        'settlementBusinessDate': !exists(json, 'settlementBusinessDate') ? undefined : (new Date(json['settlementBusinessDate'])),
        'reimbursementDateTime': !exists(json, 'reimbursementDateTime') ? undefined : (new Date(json['reimbursementDateTime'])),
        'reimbursementBusinessDate': !exists(json, 'reimbursementBusinessDate') ? undefined : (new Date(json['reimbursementBusinessDate'])),
        'fbaBillGenDate': !exists(json, 'fbaBillGenDate') ? undefined : (new Date(json['fbaBillGenDate'])),
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'flexibleBenefitAward': !exists(json, 'flexibleBenefitAward') ? undefined : json['flexibleBenefitAward'],
        'posted': !exists(json, 'posted') ? undefined : json['posted'],
        'settled': !exists(json, 'settled') ? undefined : json['settled'],
        'reimbursed': !exists(json, 'reimbursed') ? undefined : json['reimbursed'],
        'resettleAllowed': !exists(json, 'resettleAllowed') ? undefined : json['resettleAllowed'],
        'reimburseAllowed': !exists(json, 'reimburseAllowed') ? undefined : json['reimburseAllowed'],
        'orphanCertificate': !exists(json, 'orphanCertificate') ? undefined : json['orphanCertificate'],
    };
}

export function FBAInfoTypeToJSON(value?: FBAInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'status': FBAStatusTypeToJSON(value.status),
        'monetaryValue': value.monetaryValue,
        'amount': value.amount,
        'postedAmount': value.postedAmount,
        'reimbursedAmount': value.reimbursedAmount,
        'postingDateTime': value.postingDateTime === undefined ? undefined : (value.postingDateTime.toISOString().substr(0,10)),
        'postingBusinessDate': value.postingBusinessDate === undefined ? undefined : (value.postingBusinessDate.toISOString().substr(0,10)),
        'settlementDateTime': value.settlementDateTime === undefined ? undefined : (value.settlementDateTime.toISOString().substr(0,10)),
        'settlementBusinessDate': value.settlementBusinessDate === undefined ? undefined : (value.settlementBusinessDate.toISOString().substr(0,10)),
        'reimbursementDateTime': value.reimbursementDateTime === undefined ? undefined : (value.reimbursementDateTime.toISOString().substr(0,10)),
        'reimbursementBusinessDate': value.reimbursementBusinessDate === undefined ? undefined : (value.reimbursementBusinessDate.toISOString().substr(0,10)),
        'fbaBillGenDate': value.fbaBillGenDate === undefined ? undefined : (value.fbaBillGenDate.toISOString().substr(0,10)),
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'flexibleBenefitAward': value.flexibleBenefitAward,
        'posted': value.posted,
        'settled': value.settled,
        'reimbursed': value.reimbursed,
        'resettleAllowed': value.resettleAllowed,
        'reimburseAllowed': value.reimburseAllowed,
        'orphanCertificate': value.orphanCertificate,
    };
}

