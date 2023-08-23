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
import type { RateAmountBoundaryType } from './RateAmountBoundaryType';
/**
 * Rate amount boundaries for the rate plan schedule. Minimum and/or maximum rate amounts can be defined for one and/or two adults.
 * @export
 * @interface RateAmountBoundariesType
 */
export interface RateAmountBoundariesType {
    /**
     *
     * @type {RateAmountBoundaryType}
     * @memberof RateAmountBoundariesType
     */
    minimum?: RateAmountBoundaryType;
    /**
     *
     * @type {RateAmountBoundaryType}
     * @memberof RateAmountBoundariesType
     */
    maximum?: RateAmountBoundaryType;
}
/**
 * Check if a given object implements the RateAmountBoundariesType interface.
 */
export declare function instanceOfRateAmountBoundariesType(value: object): boolean;
export declare function RateAmountBoundariesTypeFromJSON(json: any): RateAmountBoundariesType;
export declare function RateAmountBoundariesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RateAmountBoundariesType;
export declare function RateAmountBoundariesTypeToJSON(value?: RateAmountBoundariesType | null): any;
