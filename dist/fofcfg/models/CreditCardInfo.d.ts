/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { ResPaymentCardType } from './ResPaymentCardType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface CreditCardInfo
 */
export interface CreditCardInfo {
    /**
     *
     * @type {ResPaymentCardType}
     * @memberof CreditCardInfo
     */
    creditCard?: ResPaymentCardType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CreditCardInfo
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CreditCardInfo
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CreditCardInfo interface.
 */
export declare function instanceOfCreditCardInfo(value: object): boolean;
export declare function CreditCardInfoFromJSON(json: any): CreditCardInfo;
export declare function CreditCardInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreditCardInfo;
export declare function CreditCardInfoToJSON(value?: CreditCardInfo | null): any;
