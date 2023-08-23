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
import type { AuthorizerCreditDetailType } from './AuthorizerCreditDetailType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Authorizer Information
 * @export
 * @interface AuthorizerCreditType
 */
export interface AuthorizerCreditType {
    /**
     *
     * @type {UniqueIDType}
     * @memberof AuthorizerCreditType
     */
    authorizerId?: UniqueIDType;
    /**
     * Application user name of the authorizer
     * @type {string}
     * @memberof AuthorizerCreditType
     */
    authorizerUserName?: string;
    /**
     * Full name of the authorizer.
     * @type {string}
     * @memberof AuthorizerCreditType
     */
    authorizerName?: string;
    /**
     * Rate code of the authorizer.
     * @type {string}
     * @memberof AuthorizerCreditType
     */
    authorizerRateCode?: string;
    /**
     * Indicates whether user has the choice to have reservation inherit rate code from the authorizer.
     * @type {boolean}
     * @memberof AuthorizerCreditType
     */
    inheritAuthorizerRateCode?: boolean;
    /**
     * Identifies the hotel code.
     * @type {string}
     * @memberof AuthorizerCreditType
     */
    hotelId?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof AuthorizerCreditType
     */
    creditLimit?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof AuthorizerCreditType
     */
    actualAmount?: CurrencyAmountType;
    /**
     * List of Comp Accounting Authorizers details
     * @type {Array<AuthorizerCreditDetailType>}
     * @memberof AuthorizerCreditType
     */
    authorizerCreditDetails?: Array<AuthorizerCreditDetailType>;
    /**
     * Transaction Date associated with the transaction.
     * @type {Date}
     * @memberof AuthorizerCreditType
     */
    transactionDate?: Date;
}
/**
 * Check if a given object implements the AuthorizerCreditType interface.
 */
export declare function instanceOfAuthorizerCreditType(value: object): boolean;
export declare function AuthorizerCreditTypeFromJSON(json: any): AuthorizerCreditType;
export declare function AuthorizerCreditTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizerCreditType;
export declare function AuthorizerCreditTypeToJSON(value?: AuthorizerCreditType | null): any;
