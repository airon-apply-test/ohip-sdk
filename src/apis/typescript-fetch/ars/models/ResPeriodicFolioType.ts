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
/**
 * Information regarding periodic folios set on the reservation.
 * @export
 * @interface ResPeriodicFolioType
 */
export interface ResPeriodicFolioType {
    /**
     * Latest date when a direct bill settlement was automatically done using the "Direct Bill Batch Folios" option.
     * @type {Date}
     * @memberof ResPeriodicFolioType
     */
    lastSettlementDate?: Date;
    /**
     * Latest date when a folio was printed using the "Periodic Batch Folios" option
     * @type {Date}
     * @memberof ResPeriodicFolioType
     */
    lastFolioDate?: Date;
    /**
     * Frequency in number of days when folios should be printed for this reservation.
     * @type {number}
     * @memberof ResPeriodicFolioType
     */
    frequency?: number;
}

/**
 * Check if a given object implements the ResPeriodicFolioType interface.
 */
export function instanceOfResPeriodicFolioType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResPeriodicFolioTypeFromJSON(json: any): ResPeriodicFolioType {
    return ResPeriodicFolioTypeFromJSONTyped(json, false);
}

export function ResPeriodicFolioTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResPeriodicFolioType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'lastSettlementDate': !exists(json, 'lastSettlementDate') ? undefined : (new Date(json['lastSettlementDate'])),
        'lastFolioDate': !exists(json, 'lastFolioDate') ? undefined : (new Date(json['lastFolioDate'])),
        'frequency': !exists(json, 'frequency') ? undefined : json['frequency'],
    };
}

export function ResPeriodicFolioTypeToJSON(value?: ResPeriodicFolioType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'lastSettlementDate': value.lastSettlementDate === undefined ? undefined : (value.lastSettlementDate.toISOString().substr(0,10)),
        'lastFolioDate': value.lastFolioDate === undefined ? undefined : (value.lastFolioDate.toISOString().substr(0,10)),
        'frequency': value.frequency,
    };
}

