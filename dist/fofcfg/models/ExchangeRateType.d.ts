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
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { CurrencyExchangeType } from './CurrencyExchangeType';
import type { StatusTypes } from './StatusTypes';
/**
 * Exchange rate configuration type.
 * @export
 * @interface ExchangeRateType
 */
export interface ExchangeRateType {
    /**
     * Property where exchange rate is set.
     * @type {string}
     * @memberof ExchangeRateType
     */
    hotelId?: string;
    /**
     * Type for exchange calculation.
     * @type {Array<CurrencyExchangeType>}
     * @memberof ExchangeRateType
     */
    exchangeType?: Array<CurrencyExchangeType>;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof ExchangeRateType
     */
    currency?: CodeDescriptionType;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof ExchangeRateType
     */
    baseCurrency?: CodeDescriptionType;
    /**
     * Date on which the rate is to become active.
     * @type {Date}
     * @memberof ExchangeRateType
     */
    beginDate?: Date;
    /**
     * Buy Rate for foreign currency.
     * @type {number}
     * @memberof ExchangeRateType
     */
    buyRate?: number;
    /**
     * Percentage of the amount that is to be a commission for the currency buy operation.
     * @type {number}
     * @memberof ExchangeRateType
     */
    buyCommission?: number;
    /**
     * Sell Rate for foreign currency.
     * @type {number}
     * @memberof ExchangeRateType
     */
    sellRate?: number;
    /**
     * Percentage of the amount that is to be a commission for the currency sell operation.
     * @type {number}
     * @memberof ExchangeRateType
     */
    sellCommission?: number;
    /**
     * Comments associated with the exchange rate.
     * @type {string}
     * @memberof ExchangeRateType
     */
    comment?: string;
    /**
     *
     * @type {StatusTypes}
     * @memberof ExchangeRateType
     */
    status?: StatusTypes;
}
/**
 * Check if a given object implements the ExchangeRateType interface.
 */
export declare function instanceOfExchangeRateType(value: object): boolean;
export declare function ExchangeRateTypeFromJSON(json: any): ExchangeRateType;
export declare function ExchangeRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExchangeRateType;
export declare function ExchangeRateTypeToJSON(value?: ExchangeRateType | null): any;
