/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * This stores the information for Bed Tax Reporting. Mainly used in Maldives.
 * @export
 * @interface BillingPrivilegesType
 */
export interface BillingPrivilegesType {
    /**
     * Flag used by interface program during check in.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    postingRestriction?: boolean;
    /**
     * Indicates if the reservation has charging privileges before arrival.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    preStayCharging?: boolean;
    /**
     * Indicates if the reservation has charging privileges after checkout.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    postStayCharging?: boolean;
    /**
     * Date when folio was closed. This works with PostStayCharging flag.
     * @type {Date}
     * @memberof BillingPrivilegesType
     */
    folioCloseDate?: Date;
    /**
     * Indicates if the guest is scheduled for automatic check out.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    scheduledCheckout?: boolean;
    /**
     * Time of automatic check out if guest is schedule for automatic check out.
     * @type {Date}
     * @memberof BillingPrivilegesType
     */
    scheduledCheckoutTime?: Date;
    /**
     * If Direct bill is authorized this will hold User ID who authorized it.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    directBillAuthorized?: boolean;
    /**
     * Indicates if the guest can do video checkout
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    videoCheckout?: boolean;
    /**
     * Indicated if a new reservation should be created and automatically checked in whenever the room is checked out. Available for pseudo room types only.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    allowAutoCheckin?: boolean;
    /**
     * Indicates if the is a candidate for auto folio settlement.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    autoSettlement?: boolean;
    /**
     * The folio settlement type for auto folio settlement.
     * @type {string}
     * @memberof BillingPrivilegesType
     */
    autoSettlementType?: string;
    /**
     * The interval of days between each auto folio settlement.
     * @type {number}
     * @memberof BillingPrivilegesType
     */
    autoSettlementFreq?: number;
    /**
     * Indicates if the reservation will be included in the Automatic Credit Limit Overages process and also be listed in the Credit Limit Overages screen results.
     * @type {boolean}
     * @memberof BillingPrivilegesType
     */
    creditLimitAutoPay?: boolean;
}

/**
 * Check if a given object implements the BillingPrivilegesType interface.
 */
export function instanceOfBillingPrivilegesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BillingPrivilegesTypeFromJSON(json: any): BillingPrivilegesType {
    return BillingPrivilegesTypeFromJSONTyped(json, false);
}

export function BillingPrivilegesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BillingPrivilegesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'postingRestriction': !exists(json, 'postingRestriction') ? undefined : json['postingRestriction'],
        'preStayCharging': !exists(json, 'preStayCharging') ? undefined : json['preStayCharging'],
        'postStayCharging': !exists(json, 'postStayCharging') ? undefined : json['postStayCharging'],
        'folioCloseDate': !exists(json, 'folioCloseDate') ? undefined : (new Date(json['folioCloseDate'])),
        'scheduledCheckout': !exists(json, 'scheduledCheckout') ? undefined : json['scheduledCheckout'],
        'scheduledCheckoutTime': !exists(json, 'scheduledCheckoutTime') ? undefined : (new Date(json['scheduledCheckoutTime'])),
        'directBillAuthorized': !exists(json, 'directBillAuthorized') ? undefined : json['directBillAuthorized'],
        'videoCheckout': !exists(json, 'videoCheckout') ? undefined : json['videoCheckout'],
        'allowAutoCheckin': !exists(json, 'allowAutoCheckin') ? undefined : json['allowAutoCheckin'],
        'autoSettlement': !exists(json, 'autoSettlement') ? undefined : json['autoSettlement'],
        'autoSettlementType': !exists(json, 'autoSettlementType') ? undefined : json['autoSettlementType'],
        'autoSettlementFreq': !exists(json, 'autoSettlementFreq') ? undefined : json['autoSettlementFreq'],
        'creditLimitAutoPay': !exists(json, 'creditLimitAutoPay') ? undefined : json['creditLimitAutoPay'],
    };
}

export function BillingPrivilegesTypeToJSON(value?: BillingPrivilegesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'postingRestriction': value.postingRestriction,
        'preStayCharging': value.preStayCharging,
        'postStayCharging': value.postStayCharging,
        'folioCloseDate': value.folioCloseDate === undefined ? undefined : (value.folioCloseDate.toISOString().substr(0,10)),
        'scheduledCheckout': value.scheduledCheckout,
        'scheduledCheckoutTime': value.scheduledCheckoutTime === undefined ? undefined : (value.scheduledCheckoutTime.toISOString().substr(0,10)),
        'directBillAuthorized': value.directBillAuthorized,
        'videoCheckout': value.videoCheckout,
        'allowAutoCheckin': value.allowAutoCheckin,
        'autoSettlement': value.autoSettlement,
        'autoSettlementType': value.autoSettlementType,
        'autoSettlementFreq': value.autoSettlementFreq,
        'creditLimitAutoPay': value.creditLimitAutoPay,
    };
}

