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
import type { InstanceLink } from './InstanceLink';
import type { PolicyApplyCriteriaType } from './PolicyApplyCriteriaType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface DetachDepositPolicyFromPaymentRequest
 */
export interface DetachDepositPolicyFromPaymentRequest {
    /**
     *
     * @type {PolicyApplyCriteriaType}
     * @memberof DetachDepositPolicyFromPaymentRequest
     */
    criteria?: PolicyApplyCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DetachDepositPolicyFromPaymentRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DetachDepositPolicyFromPaymentRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DetachDepositPolicyFromPaymentRequest interface.
 */
export declare function instanceOfDetachDepositPolicyFromPaymentRequest(value: object): boolean;
export declare function DetachDepositPolicyFromPaymentRequestFromJSON(json: any): DetachDepositPolicyFromPaymentRequest;
export declare function DetachDepositPolicyFromPaymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DetachDepositPolicyFromPaymentRequest;
export declare function DetachDepositPolicyFromPaymentRequestToJSON(value?: DetachDepositPolicyFromPaymentRequest | null): any;
