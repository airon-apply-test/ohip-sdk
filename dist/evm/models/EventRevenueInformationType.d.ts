/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { RevenueTypeType } from './RevenueTypeType';
/**
 * Pertain event's revenue information.
 * @export
 * @interface EventRevenueInformationType
 */
export interface EventRevenueInformationType {
    /**
     *
     * @type {RevenueTypeType}
     * @memberof EventRevenueInformationType
     */
    revenueType?: RevenueTypeType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    forecastRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    expectedRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    guaranteedRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    onTheBooksRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    billedRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    actualRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof EventRevenueInformationType
     */
    totalCost?: CurrencyAmountType;
    /**
     * Flag to indicate forecast revenue is a flat amount, therefore NOT per person.
     * @type {boolean}
     * @memberof EventRevenueInformationType
     */
    flatForecastRevenue?: boolean;
    /**
     * Flag indicating revenues are from a custom revenue type.
     * @type {boolean}
     * @memberof EventRevenueInformationType
     */
    custom?: boolean;
    /**
     * Flag indicating to ignore forecast figures.
     * @type {boolean}
     * @memberof EventRevenueInformationType
     */
    ignoreForecast?: boolean;
    /**
     * Flag indicating if the revenues are generated by a catering package.
     * @type {boolean}
     * @memberof EventRevenueInformationType
     */
    packaged?: boolean;
    /**
     * Sort order for revenues.
     * @type {number}
     * @memberof EventRevenueInformationType
     */
    orderBy?: number;
}
/**
 * Check if a given object implements the EventRevenueInformationType interface.
 */
export declare function instanceOfEventRevenueInformationType(value: object): boolean;
export declare function EventRevenueInformationTypeFromJSON(json: any): EventRevenueInformationType;
export declare function EventRevenueInformationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EventRevenueInformationType;
export declare function EventRevenueInformationTypeToJSON(value?: EventRevenueInformationType | null): any;
