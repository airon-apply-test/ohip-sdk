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
/**
 * Advanced Dynamic Base Rate type
 * @export
 * @interface AdvancedDynamicBaseRateType
 */
export interface AdvancedDynamicBaseRateType {
    /**
     * Rate plan code used to Advanced Dynamically base the rate on.
     * @type {string}
     * @memberof AdvancedDynamicBaseRateType
     */
    basedOnRatePlan?: string;
    /**
     * Rounding style used for the calculated rate amounts. Valid values are U,D,N,C,F which means Up, Down, None, Up-Keep Decimal, Down - Keep Decimal.
     * @type {string}
     * @memberof AdvancedDynamicBaseRateType
     */
    rounding?: string;
    /**
     * While showing availability, do system need to compare the static rates defined for the rate with Advanced dynamically computed base rate pricing? If true, availability shows lower of the these two pricing. If false, availability shows the advanced dynamically calculated pricing.
     * @type {boolean}
     * @memberof AdvancedDynamicBaseRateType
     */
    compareWithRateSchedules?: boolean;
    /**
     * Rate Plan code.
     * @type {Array<string>}
     * @memberof AdvancedDynamicBaseRateType
     */
    advancedDependentRatePlans?: Array<string>;
}
/**
 * Check if a given object implements the AdvancedDynamicBaseRateType interface.
 */
export declare function instanceOfAdvancedDynamicBaseRateType(value: object): boolean;
export declare function AdvancedDynamicBaseRateTypeFromJSON(json: any): AdvancedDynamicBaseRateType;
export declare function AdvancedDynamicBaseRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AdvancedDynamicBaseRateType;
export declare function AdvancedDynamicBaseRateTypeToJSON(value?: AdvancedDynamicBaseRateType | null): any;
