/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Defines Child's Age.
 * @export
 * @interface ChildAgeType
 */
export interface ChildAgeType {
    /**
     * Age of a child in years.
     * @type {number}
     * @memberof ChildAgeType
     */
    age?: number;
}
/**
 * Check if a given object implements the ChildAgeType interface.
 */
export declare function instanceOfChildAgeType(value: object): boolean;
export declare function ChildAgeTypeFromJSON(json: any): ChildAgeType;
export declare function ChildAgeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChildAgeType;
export declare function ChildAgeTypeToJSON(value?: ChildAgeType | null): any;
