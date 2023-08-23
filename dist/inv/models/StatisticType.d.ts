/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { StatisticCodeType } from './StatisticCodeType';
/**
 * Defines all details needed to create a statistical report.
 * @export
 * @interface StatisticType
 */
export interface StatisticType {
    /**
     * Statistic Codes.
     * @type {Array<StatisticCodeType>}
     * @memberof StatisticType
     */
    statistics?: Array<StatisticCodeType>;
    /**
     * A text field used to communicate the proper name of the hotel.
     * @type {string}
     * @memberof StatisticType
     */
    hotelName?: string;
    /**
     * Identifies the type of statistics collected. Each ReportCode corresponds to a set of category summaries based upon a predetermined agreement.
     * @type {string}
     * @memberof StatisticType
     */
    reportCode?: string;
    /**
     * This element has revenue amount data for its revenue category such as Room Revenue, Food and Beverage Revenue.
     * @type {string}
     * @memberof StatisticType
     */
    description?: string;
}
/**
 * Check if a given object implements the StatisticType interface.
 */
export declare function instanceOfStatisticType(value: object): boolean;
export declare function StatisticTypeFromJSON(json: any): StatisticType;
export declare function StatisticTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StatisticType;
export declare function StatisticTypeToJSON(value?: StatisticType | null): any;
