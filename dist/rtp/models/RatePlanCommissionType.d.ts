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
 *
 * @export
 * @interface RatePlanCommissionType
 */
export interface RatePlanCommissionType {
    /**
     * commission code used by the rate plan.
     * @type {string}
     * @memberof RatePlanCommissionType
     */
    commissionCode?: string;
    /**
     * commission percentage used by the rate plan.
     * @type {number}
     * @memberof RatePlanCommissionType
     */
    commissionPercentage?: number;
    /**
     * commission amount used by the rate plan.
     * @type {number}
     * @memberof RatePlanCommissionType
     */
    commissionAmount?: number;
}
/**
 * Check if a given object implements the RatePlanCommissionType interface.
 */
export declare function instanceOfRatePlanCommissionType(value: object): boolean;
export declare function RatePlanCommissionTypeFromJSON(json: any): RatePlanCommissionType;
export declare function RatePlanCommissionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RatePlanCommissionType;
export declare function RatePlanCommissionTypeToJSON(value?: RatePlanCommissionType | null): any;
