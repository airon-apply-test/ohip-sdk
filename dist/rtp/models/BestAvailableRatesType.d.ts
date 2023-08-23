/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BestAvailableRatesTypeDuration } from './BestAvailableRatesTypeDuration';
import type { BestAvailableRatesTypeLengthOfStay } from './BestAvailableRatesTypeLengthOfStay';
import type { DateRangeType } from './DateRangeType';
/**
 * The best available rate for rate code(s)
 * @export
 * @interface BestAvailableRatesType
 */
export interface BestAvailableRatesType {
    /**
     * Hotel Code for Best Available Rate
     * @type {string}
     * @memberof BestAvailableRatesType
     */
    hotelId?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof BestAvailableRatesType
     */
    rateCode?: Array<string>;
    /**
     *
     * @type {DateRangeType}
     * @memberof BestAvailableRatesType
     */
    dateRange?: DateRangeType;
    /**
     *
     * @type {BestAvailableRatesTypeDuration}
     * @memberof BestAvailableRatesType
     */
    duration?: BestAvailableRatesTypeDuration;
    /**
     *
     * @type {BestAvailableRatesTypeLengthOfStay}
     * @memberof BestAvailableRatesType
     */
    lengthOfStay?: BestAvailableRatesTypeLengthOfStay;
}
/**
 * Check if a given object implements the BestAvailableRatesType interface.
 */
export declare function instanceOfBestAvailableRatesType(value: object): boolean;
export declare function BestAvailableRatesTypeFromJSON(json: any): BestAvailableRatesType;
export declare function BestAvailableRatesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BestAvailableRatesType;
export declare function BestAvailableRatesTypeToJSON(value?: BestAvailableRatesType | null): any;
