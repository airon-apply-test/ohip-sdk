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
import type { BedTaxReportingType } from './BedTaxReportingType';
import {
    BedTaxReportingTypeFromJSON,
    BedTaxReportingTypeFromJSONTyped,
    BedTaxReportingTypeToJSON,
} from './BedTaxReportingType';
import type { BillingPrivilegesType } from './BillingPrivilegesType';
import {
    BillingPrivilegesTypeFromJSON,
    BillingPrivilegesTypeFromJSONTyped,
    BillingPrivilegesTypeToJSON,
} from './BillingPrivilegesType';
import type { FolioTextsTypeInner } from './FolioTextsTypeInner';
import {
    FolioTextsTypeInnerFromJSON,
    FolioTextsTypeInnerFromJSONTyped,
    FolioTextsTypeInnerToJSON,
} from './FolioTextsTypeInner';
import type { ResCompAccountingType } from './ResCompAccountingType';
import {
    ResCompAccountingTypeFromJSON,
    ResCompAccountingTypeFromJSONTyped,
    ResCompAccountingTypeToJSON,
} from './ResCompAccountingType';
import type { ResPeriodicFolioType } from './ResPeriodicFolioType';
import {
    ResPeriodicFolioTypeFromJSON,
    ResPeriodicFolioTypeFromJSONTyped,
    ResPeriodicFolioTypeToJSON,
} from './ResPeriodicFolioType';
import type { ResPreConfiguredRoutingInstrType } from './ResPreConfiguredRoutingInstrType';
import {
    ResPreConfiguredRoutingInstrTypeFromJSON,
    ResPreConfiguredRoutingInstrTypeFromJSONTyped,
    ResPreConfiguredRoutingInstrTypeToJSON,
} from './ResPreConfiguredRoutingInstrType';
import type { ResRevenueBalanceType } from './ResRevenueBalanceType';
import {
    ResRevenueBalanceTypeFromJSON,
    ResRevenueBalanceTypeFromJSONTyped,
    ResRevenueBalanceTypeToJSON,
} from './ResRevenueBalanceType';
import type { ReservationTaxTypeInfo } from './ReservationTaxTypeInfo';
import {
    ReservationTaxTypeInfoFromJSON,
    ReservationTaxTypeInfoFromJSONTyped,
    ReservationTaxTypeInfoToJSON,
} from './ReservationTaxTypeInfo';

/**
 * Cashiering Information for the reservation.
 * @export
 * @interface ResCashieringType
 */
export interface ResCashieringType {
    /**
     * 
     * @type {ResRevenueBalanceType}
     * @memberof ResCashieringType
     */
    revenuesAndBalances?: ResRevenueBalanceType;
    /**
     * 
     * @type {BillingPrivilegesType}
     * @memberof ResCashieringType
     */
    billingPrivileges?: BillingPrivilegesType;
    /**
     * 
     * @type {ReservationTaxTypeInfo}
     * @memberof ResCashieringType
     */
    taxType?: ReservationTaxTypeInfo;
    /**
     * 
     * @type {BedTaxReportingType}
     * @memberof ResCashieringType
     */
    bedTaxReporting?: BedTaxReportingType;
    /**
     * This stores the description for the type of tax calculation especially with tax exemption, etc.
     * @type {Array<FolioTextsTypeInner>}
     * @memberof ResCashieringType
     */
    folioTexts?: Array<FolioTextsTypeInner>;
    /**
     * 
     * @type {ResPeriodicFolioType}
     * @memberof ResCashieringType
     */
    periodicFolio?: ResPeriodicFolioType;
    /**
     * 
     * @type {ResCompAccountingType}
     * @memberof ResCashieringType
     */
    compAccounting?: ResCompAccountingType;
    /**
     * 
     * @type {ResPreConfiguredRoutingInstrType}
     * @memberof ResCashieringType
     */
    reservationPreConfiguredRoutingInstruction?: ResPreConfiguredRoutingInstrType;
    /**
     * The guest from whom payment has to be recovered (direct guest).
     * @type {boolean}
     * @memberof ResCashieringType
     */
    financiallyResponsible?: boolean;
    /**
     * In case of Appartment style billing indicates whether a prorated amount should be used for an Apartment Style Billing rate.
     * @type {boolean}
     * @memberof ResCashieringType
     */
    proratedBilling?: boolean;
    /**
     * Date of the last Room And Tax posting. Used primarily to know the date in case of Advance Billing.
     * @type {Date}
     * @memberof ResCashieringType
     */
    lastRoomAndTaxPostedDate?: Date;
    /**
     * This attribute is to verify if reverse check-in is allowed for the reservation.
     * @type {boolean}
     * @memberof ResCashieringType
     */
    reverseCheckInAllowed?: boolean;
    /**
     * This attribute is to verify if reverse advance check-in is allowed for the reservation.
     * @type {boolean}
     * @memberof ResCashieringType
     */
    reverseAdvanceCheckInAllowed?: boolean;
    /**
     * Specifies whether reservation has a financial transaction associated with it.
     * @type {boolean}
     * @memberof ResCashieringType
     */
    transactionsPosted?: boolean;
}

/**
 * Check if a given object implements the ResCashieringType interface.
 */
export function instanceOfResCashieringType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResCashieringTypeFromJSON(json: any): ResCashieringType {
    return ResCashieringTypeFromJSONTyped(json, false);
}

export function ResCashieringTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResCashieringType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'revenuesAndBalances': !exists(json, 'revenuesAndBalances') ? undefined : ResRevenueBalanceTypeFromJSON(json['revenuesAndBalances']),
        'billingPrivileges': !exists(json, 'billingPrivileges') ? undefined : BillingPrivilegesTypeFromJSON(json['billingPrivileges']),
        'taxType': !exists(json, 'taxType') ? undefined : ReservationTaxTypeInfoFromJSON(json['taxType']),
        'bedTaxReporting': !exists(json, 'bedTaxReporting') ? undefined : BedTaxReportingTypeFromJSON(json['bedTaxReporting']),
        'folioTexts': !exists(json, 'folioTexts') ? undefined : ((json['folioTexts'] as Array<any>).map(FolioTextsTypeInnerFromJSON)),
        'periodicFolio': !exists(json, 'periodicFolio') ? undefined : ResPeriodicFolioTypeFromJSON(json['periodicFolio']),
        'compAccounting': !exists(json, 'compAccounting') ? undefined : ResCompAccountingTypeFromJSON(json['compAccounting']),
        'reservationPreConfiguredRoutingInstruction': !exists(json, 'reservationPreConfiguredRoutingInstruction') ? undefined : ResPreConfiguredRoutingInstrTypeFromJSON(json['reservationPreConfiguredRoutingInstruction']),
        'financiallyResponsible': !exists(json, 'financiallyResponsible') ? undefined : json['financiallyResponsible'],
        'proratedBilling': !exists(json, 'proratedBilling') ? undefined : json['proratedBilling'],
        'lastRoomAndTaxPostedDate': !exists(json, 'lastRoomAndTaxPostedDate') ? undefined : (new Date(json['lastRoomAndTaxPostedDate'])),
        'reverseCheckInAllowed': !exists(json, 'reverseCheckInAllowed') ? undefined : json['reverseCheckInAllowed'],
        'reverseAdvanceCheckInAllowed': !exists(json, 'reverseAdvanceCheckInAllowed') ? undefined : json['reverseAdvanceCheckInAllowed'],
        'transactionsPosted': !exists(json, 'transactionsPosted') ? undefined : json['transactionsPosted'],
    };
}

export function ResCashieringTypeToJSON(value?: ResCashieringType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'revenuesAndBalances': ResRevenueBalanceTypeToJSON(value.revenuesAndBalances),
        'billingPrivileges': BillingPrivilegesTypeToJSON(value.billingPrivileges),
        'taxType': ReservationTaxTypeInfoToJSON(value.taxType),
        'bedTaxReporting': BedTaxReportingTypeToJSON(value.bedTaxReporting),
        'folioTexts': value.folioTexts === undefined ? undefined : ((value.folioTexts as Array<any>).map(FolioTextsTypeInnerToJSON)),
        'periodicFolio': ResPeriodicFolioTypeToJSON(value.periodicFolio),
        'compAccounting': ResCompAccountingTypeToJSON(value.compAccounting),
        'reservationPreConfiguredRoutingInstruction': ResPreConfiguredRoutingInstrTypeToJSON(value.reservationPreConfiguredRoutingInstruction),
        'financiallyResponsible': value.financiallyResponsible,
        'proratedBilling': value.proratedBilling,
        'lastRoomAndTaxPostedDate': value.lastRoomAndTaxPostedDate === undefined ? undefined : (value.lastRoomAndTaxPostedDate.toISOString().substr(0,10)),
        'reverseCheckInAllowed': value.reverseCheckInAllowed,
        'reverseAdvanceCheckInAllowed': value.reverseAdvanceCheckInAllowed,
        'transactionsPosted': value.transactionsPosted,
    };
}

