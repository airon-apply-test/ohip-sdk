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
import type { CashierClosureCurrencyAmountType } from './CashierClosureCurrencyAmountType';
import type { CashierClosureNoType } from './CashierClosureNoType';
import type { HotelCashierClosureCurrencyAmountType } from './HotelCashierClosureCurrencyAmountType';
/**
 * Cashier Shift Totals and Per Hotel amounts group by currency code
 * @export
 * @interface CashierClosureType
 */
export interface CashierClosureType {
    /**
     *
     * @type {CashierClosureNoType}
     * @memberof CashierClosureType
     */
    cashierInfo?: CashierClosureNoType;
    /**
     * Currency Total Amounts for Cashier Shift.
     * @type {Array<CashierClosureCurrencyAmountType>}
     * @memberof CashierClosureType
     */
    summaryCurrencyAmountList?: Array<CashierClosureCurrencyAmountType>;
    /**
     * Currency Hotels Amounts for Cashier Shift.
     * @type {Array<HotelCashierClosureCurrencyAmountType>}
     * @memberof CashierClosureType
     */
    hotelCurrencyAmountList?: Array<HotelCashierClosureCurrencyAmountType>;
}
/**
 * Check if a given object implements the CashierClosureType interface.
 */
export declare function instanceOfCashierClosureType(value: object): boolean;
export declare function CashierClosureTypeFromJSON(json: any): CashierClosureType;
export declare function CashierClosureTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierClosureType;
export declare function CashierClosureTypeToJSON(value?: CashierClosureType | null): any;
