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
import type { ARAccountCriteriaType } from './ARAccountCriteriaType';
import {
    ARAccountCriteriaTypeFromJSON,
    ARAccountCriteriaTypeFromJSONTyped,
    ARAccountCriteriaTypeToJSON,
} from './ARAccountCriteriaType';
import type { ChargeCriteriaType } from './ChargeCriteriaType';
import {
    ChargeCriteriaTypeFromJSON,
    ChargeCriteriaTypeFromJSONTyped,
    ChargeCriteriaTypeToJSON,
} from './ChargeCriteriaType';
import type { FiscalServiceType } from './FiscalServiceType';
import {
    FiscalServiceTypeFromJSON,
    FiscalServiceTypeFromJSONTyped,
    FiscalServiceTypeToJSON,
} from './FiscalServiceType';
import type { NameValueHeaderDetailType } from './NameValueHeaderDetailType';
import {
    NameValueHeaderDetailTypeFromJSON,
    NameValueHeaderDetailTypeFromJSONTyped,
    NameValueHeaderDetailTypeToJSON,
} from './NameValueHeaderDetailType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Criteria to Create a new AR Invoice.
 * @export
 * @interface CreateARInvoiceCriteriaType
 */
export interface CreateARInvoiceCriteriaType {
    /**
     * Custom Folio Name Value Information to be saved
     * @type {Array<NameValueHeaderDetailType>}
     * @memberof CreateARInvoiceCriteriaType
     */
    folioNameValue?: Array<NameValueHeaderDetailType>;
    /**
     * 
     * @type {ARAccountCriteriaType}
     * @memberof CreateARInvoiceCriteriaType
     */
    account?: ARAccountCriteriaType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof CreateARInvoiceCriteriaType
     */
    guestProfileId?: UniqueIDType;
    /**
     * Collection of Charges to be posted.
     * @type {Array<ChargeCriteriaType>}
     * @memberof CreateARInvoiceCriteriaType
     */
    charges?: Array<ChargeCriteriaType>;
    /**
     * User-defined invoice reference.
     * @type {string}
     * @memberof CreateARInvoiceCriteriaType
     */
    reference?: string;
    /**
     * User-defined invoice remark.
     * @type {string}
     * @memberof CreateARInvoiceCriteriaType
     */
    remark?: string;
    /**
     * Invoice market code.
     * @type {string}
     * @memberof CreateARInvoiceCriteriaType
     */
    market?: string;
    /**
     * Invoice room class code.
     * @type {string}
     * @memberof CreateARInvoiceCriteriaType
     */
    roomClass?: string;
    /**
     * Invoice source code.
     * @type {string}
     * @memberof CreateARInvoiceCriteriaType
     */
    source?: string;
    /**
     * 
     * @type {FiscalServiceType}
     * @memberof CreateARInvoiceCriteriaType
     */
    fiscalFolioInfo?: FiscalServiceType;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof CreateARInvoiceCriteriaType
     */
    cashierId?: number;
    /**
     * 
     * @type {boolean}
     * @memberof CreateARInvoiceCriteriaType
     */
    overrideCreditHoldCheck?: boolean;
}

/**
 * Check if a given object implements the CreateARInvoiceCriteriaType interface.
 */
export function instanceOfCreateARInvoiceCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateARInvoiceCriteriaTypeFromJSON(json: any): CreateARInvoiceCriteriaType {
    return CreateARInvoiceCriteriaTypeFromJSONTyped(json, false);
}

export function CreateARInvoiceCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateARInvoiceCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'folioNameValue': !exists(json, 'folioNameValue') ? undefined : ((json['folioNameValue'] as Array<any>).map(NameValueHeaderDetailTypeFromJSON)),
        'account': !exists(json, 'account') ? undefined : ARAccountCriteriaTypeFromJSON(json['account']),
        'guestProfileId': !exists(json, 'guestProfileId') ? undefined : UniqueIDTypeFromJSON(json['guestProfileId']),
        'charges': !exists(json, 'charges') ? undefined : ((json['charges'] as Array<any>).map(ChargeCriteriaTypeFromJSON)),
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'remark': !exists(json, 'remark') ? undefined : json['remark'],
        'market': !exists(json, 'market') ? undefined : json['market'],
        'roomClass': !exists(json, 'roomClass') ? undefined : json['roomClass'],
        'source': !exists(json, 'source') ? undefined : json['source'],
        'fiscalFolioInfo': !exists(json, 'fiscalFolioInfo') ? undefined : FiscalServiceTypeFromJSON(json['fiscalFolioInfo']),
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
        'overrideCreditHoldCheck': !exists(json, 'overrideCreditHoldCheck') ? undefined : json['overrideCreditHoldCheck'],
    };
}

export function CreateARInvoiceCriteriaTypeToJSON(value?: CreateARInvoiceCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'folioNameValue': value.folioNameValue === undefined ? undefined : ((value.folioNameValue as Array<any>).map(NameValueHeaderDetailTypeToJSON)),
        'account': ARAccountCriteriaTypeToJSON(value.account),
        'guestProfileId': UniqueIDTypeToJSON(value.guestProfileId),
        'charges': value.charges === undefined ? undefined : ((value.charges as Array<any>).map(ChargeCriteriaTypeToJSON)),
        'reference': value.reference,
        'remark': value.remark,
        'market': value.market,
        'roomClass': value.roomClass,
        'source': value.source,
        'fiscalFolioInfo': FiscalServiceTypeToJSON(value.fiscalFolioInfo),
        'cashierId': value.cashierId,
        'overrideCreditHoldCheck': value.overrideCreditHoldCheck,
    };
}

