/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Search match indicating attribute and the matching value.
 * @export
 * @interface SearchMatchType
 */
export interface SearchMatchType {
    /**
     * Search match attribute.
     * @type {string}
     * @memberof SearchMatchType
     */
    attribute?: string;
    /**
     * Search match value.
     * @type {string}
     * @memberof SearchMatchType
     */
    value?: string;
}
/**
 * Check if a given object implements the SearchMatchType interface.
 */
export declare function instanceOfSearchMatchType(value: object): boolean;
export declare function SearchMatchTypeFromJSON(json: any): SearchMatchType;
export declare function SearchMatchTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchMatchType;
export declare function SearchMatchTypeToJSON(value?: SearchMatchType | null): any;
