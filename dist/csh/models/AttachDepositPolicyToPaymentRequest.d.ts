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
 * @interface AttachDepositPolicyToPaymentRequest
 */
export interface AttachDepositPolicyToPaymentRequest {
    /**
     *
     * @type {PolicyApplyCriteriaType}
     * @memberof AttachDepositPolicyToPaymentRequest
     */
    criteria?: PolicyApplyCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof AttachDepositPolicyToPaymentRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof AttachDepositPolicyToPaymentRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the AttachDepositPolicyToPaymentRequest interface.
 */
export declare function instanceOfAttachDepositPolicyToPaymentRequest(value: object): boolean;
export declare function AttachDepositPolicyToPaymentRequestFromJSON(json: any): AttachDepositPolicyToPaymentRequest;
export declare function AttachDepositPolicyToPaymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AttachDepositPolicyToPaymentRequest;
export declare function AttachDepositPolicyToPaymentRequestToJSON(value?: AttachDepositPolicyToPaymentRequest | null): any;
