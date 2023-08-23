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
 * Report History Type used as based type for Remiders and Statements History types.
 * @export
 * @interface ARStatementHistoryType
 */
export interface ARStatementHistoryType {
    /**
     * Report Name.
     * @type {string}
     * @memberof ARStatementHistoryType
     */
    reportName?: string;
    /**
     * Report file name when exists to allow report re-printing.
     * @type {string}
     * @memberof ARStatementHistoryType
     */
    reportFileName?: string;
    /**
     * The Reminder Letter name which is to be used for this Reminder based on the setup on the Account Type.
     * @type {Date}
     * @memberof ARStatementHistoryType
     */
    dateSent?: Date;
    /**
     * When using Statement Numbering, a unique number is associated to the Statement.
     * @type {number}
     * @memberof ARStatementHistoryType
     */
    statementNo?: number;
}

/**
 * Check if a given object implements the ARStatementHistoryType interface.
 */
export function instanceOfARStatementHistoryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ARStatementHistoryTypeFromJSON(json: any): ARStatementHistoryType {
    return ARStatementHistoryTypeFromJSONTyped(json, false);
}

export function ARStatementHistoryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARStatementHistoryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reportName': !exists(json, 'reportName') ? undefined : json['reportName'],
        'reportFileName': !exists(json, 'reportFileName') ? undefined : json['reportFileName'],
        'dateSent': !exists(json, 'dateSent') ? undefined : (new Date(json['dateSent'])),
        'statementNo': !exists(json, 'statementNo') ? undefined : json['statementNo'],
    };
}

export function ARStatementHistoryTypeToJSON(value?: ARStatementHistoryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reportName': value.reportName,
        'reportFileName': value.reportFileName,
        'dateSent': value.dateSent === undefined ? undefined : (value.dateSent.toISOString().substr(0,10)),
        'statementNo': value.statementNo,
    };
}

