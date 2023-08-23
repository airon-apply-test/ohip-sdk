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
import type { ARStatementHistoryType } from './ARStatementHistoryType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Statement History.
 * @export
 * @interface StatementHistory
 */
export interface StatementHistory {
    /**
     *
     * @type {Array<ARStatementHistoryType>}
     * @memberof StatementHistory
     */
    aRStatementHistory?: Array<ARStatementHistoryType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof StatementHistory
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof StatementHistory
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the StatementHistory interface.
 */
export declare function instanceOfStatementHistory(value: object): boolean;
export declare function StatementHistoryFromJSON(json: any): StatementHistory;
export declare function StatementHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): StatementHistory;
export declare function StatementHistoryToJSON(value?: StatementHistory | null): any;
