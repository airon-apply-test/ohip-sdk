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
 *
 * @export
 * @interface ReverseFolioTaxPaymentRequest
 */
export interface ReverseFolioTaxPaymentRequest {
    /**
     *
     * @type {DeleteTransactionCriteriaType}
     * @memberof ReverseFolioTaxPaymentRequest
     */
    reverseCriteria?: DeleteTransactionCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ReverseFolioTaxPaymentRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReverseFolioTaxPaymentRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ReverseFolioTaxPaymentRequest interface.
 */
export declare function instanceOfReverseFolioTaxPaymentRequest(value: object): boolean;
export declare function ReverseFolioTaxPaymentRequestFromJSON(json: any): ReverseFolioTaxPaymentRequest;
export declare function ReverseFolioTaxPaymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReverseFolioTaxPaymentRequest;
export declare function ReverseFolioTaxPaymentRequestToJSON(value?: ReverseFolioTaxPaymentRequest | null): any;
