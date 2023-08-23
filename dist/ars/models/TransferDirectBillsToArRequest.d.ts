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
import type { InstanceLink } from './InstanceLink';
import type { TransferDirectBillsToArCriteriaType } from './TransferDirectBillsToArCriteriaType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface TransferDirectBillsToArRequest
 */
export interface TransferDirectBillsToArRequest {
    /**
     *
     * @type {TransferDirectBillsToArCriteriaType}
     * @memberof TransferDirectBillsToArRequest
     */
    criteria?: TransferDirectBillsToArCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof TransferDirectBillsToArRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TransferDirectBillsToArRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the TransferDirectBillsToArRequest interface.
 */
export declare function instanceOfTransferDirectBillsToArRequest(value: object): boolean;
export declare function TransferDirectBillsToArRequestFromJSON(json: any): TransferDirectBillsToArRequest;
export declare function TransferDirectBillsToArRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransferDirectBillsToArRequest;
export declare function TransferDirectBillsToArRequestToJSON(value?: TransferDirectBillsToArRequest | null): any;
