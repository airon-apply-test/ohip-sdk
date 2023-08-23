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
import type { FolioStatusType } from './FolioStatusType';
import {
    FolioStatusTypeFromJSON,
    FolioStatusTypeFromJSONTyped,
    FolioStatusTypeToJSON,
} from './FolioStatusType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Information about comp folio that was auto settled.
 * @export
 * @interface AutoSettleCompFolioType
 */
export interface AutoSettleCompFolioType {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof AutoSettleCompFolioType
     */
    payeeId?: UniqueIDType;
    /**
     * Name of the payee.
     * @type {string}
     * @memberof AutoSettleCompFolioType
     */
    payeeName?: string;
    /**
     * Comp folio window that was attempted to be auto settled.
     * @type {number}
     * @memberof AutoSettleCompFolioType
     */
    folioView?: number;
    /**
     * Folio Number.
     * @type {number}
     * @memberof AutoSettleCompFolioType
     */
    folioNo?: number;
    /**
     * Invoice No after the folio is generated. Same invoice number may be referred in multiple folios
     * @type {number}
     * @memberof AutoSettleCompFolioType
     */
    invoiceNo?: number;
    /**
     * The Fiscal Bill number of this posting
     * @type {string}
     * @memberof AutoSettleCompFolioType
     */
    fiscalBillNo?: string;
    /**
     * The name of the Folio Type used for the Folio Number sequence.
     * @type {string}
     * @memberof AutoSettleCompFolioType
     */
    folioTypeName?: string;
    /**
     * Internal window ID which is unique to the reservation. This ID can only be used for reference.
     * @type {string}
     * @memberof AutoSettleCompFolioType
     */
    internalFolioWindowID?: string;
    /**
     * Date of Folio Generation.
     * @type {Date}
     * @memberof AutoSettleCompFolioType
     */
    folioDate?: Date;
    /**
     * 
     * @type {FolioStatusType}
     * @memberof AutoSettleCompFolioType
     */
    folioStatus?: FolioStatusType;
    /**
     * The folio number with prefix value.
     * @type {string}
     * @memberof AutoSettleCompFolioType
     */
    folioNoWithPrefix?: string;
    /**
     * Unique sequence number. Used to identify the current folio tax record.
     * @type {number}
     * @memberof AutoSettleCompFolioType
     */
    folioSeqNo?: number;
}

/**
 * Check if a given object implements the AutoSettleCompFolioType interface.
 */
export function instanceOfAutoSettleCompFolioType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AutoSettleCompFolioTypeFromJSON(json: any): AutoSettleCompFolioType {
    return AutoSettleCompFolioTypeFromJSONTyped(json, false);
}

export function AutoSettleCompFolioTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AutoSettleCompFolioType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'payeeId': !exists(json, 'payeeId') ? undefined : UniqueIDTypeFromJSON(json['payeeId']),
        'payeeName': !exists(json, 'payeeName') ? undefined : json['payeeName'],
        'folioView': !exists(json, 'folioView') ? undefined : json['folioView'],
        'folioNo': !exists(json, 'folioNo') ? undefined : json['folioNo'],
        'invoiceNo': !exists(json, 'invoiceNo') ? undefined : json['invoiceNo'],
        'fiscalBillNo': !exists(json, 'fiscalBillNo') ? undefined : json['fiscalBillNo'],
        'folioTypeName': !exists(json, 'folioTypeName') ? undefined : json['folioTypeName'],
        'internalFolioWindowID': !exists(json, 'internalFolioWindowID') ? undefined : json['internalFolioWindowID'],
        'folioDate': !exists(json, 'folioDate') ? undefined : (new Date(json['folioDate'])),
        'folioStatus': !exists(json, 'folioStatus') ? undefined : FolioStatusTypeFromJSON(json['folioStatus']),
        'folioNoWithPrefix': !exists(json, 'folioNoWithPrefix') ? undefined : json['folioNoWithPrefix'],
        'folioSeqNo': !exists(json, 'folioSeqNo') ? undefined : json['folioSeqNo'],
    };
}

export function AutoSettleCompFolioTypeToJSON(value?: AutoSettleCompFolioType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'payeeId': UniqueIDTypeToJSON(value.payeeId),
        'payeeName': value.payeeName,
        'folioView': value.folioView,
        'folioNo': value.folioNo,
        'invoiceNo': value.invoiceNo,
        'fiscalBillNo': value.fiscalBillNo,
        'folioTypeName': value.folioTypeName,
        'internalFolioWindowID': value.internalFolioWindowID,
        'folioDate': value.folioDate === undefined ? undefined : (value.folioDate.toISOString().substr(0,10)),
        'folioStatus': FolioStatusTypeToJSON(value.folioStatus),
        'folioNoWithPrefix': value.folioNoWithPrefix,
        'folioSeqNo': value.folioSeqNo,
    };
}

