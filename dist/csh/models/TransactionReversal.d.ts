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
import type { DeleteTransactionCriteriaType } from './DeleteTransactionCriteriaType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to reverse a Payment Tax record for Thailand Tax requirement.
 * @export
 * @interface TransactionReversal
 */
export interface TransactionReversal {
    /**
     *
     * @type {DeleteTransactionCriteriaType}
     * @memberof TransactionReversal
     */
    reverseCriteria?: DeleteTransactionCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof TransactionReversal
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TransactionReversal
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the TransactionReversal interface.
 */
export declare function instanceOfTransactionReversal(value: object): boolean;
export declare function TransactionReversalFromJSON(json: any): TransactionReversal;
export declare function TransactionReversalFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransactionReversal;
export declare function TransactionReversalToJSON(value?: TransactionReversal | null): any;
