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
import type { RatePlanShortInfoType } from './RatePlanShortInfoType';
/**
 * List of summary information for the rate plan code.
 * @export
 * @interface RatePlansSummaryRatePlanShortInfoList
 */
export interface RatePlansSummaryRatePlanShortInfoList {
    /**
     * List of summary information for the rate plan code.
     * @type {Array<RatePlanShortInfoType>}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    ratePlanShortInfo?: Array<RatePlanShortInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    count?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    limit?: number;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof RatePlansSummaryRatePlanShortInfoList
     */
    totalPages?: number;
}
/**
 * Check if a given object implements the RatePlansSummaryRatePlanShortInfoList interface.
 */
export declare function instanceOfRatePlansSummaryRatePlanShortInfoList(value: object): boolean;
export declare function RatePlansSummaryRatePlanShortInfoListFromJSON(json: any): RatePlansSummaryRatePlanShortInfoList;
export declare function RatePlansSummaryRatePlanShortInfoListFromJSONTyped(json: any, ignoreDiscriminator: boolean): RatePlansSummaryRatePlanShortInfoList;
export declare function RatePlansSummaryRatePlanShortInfoListToJSON(value?: RatePlansSummaryRatePlanShortInfoList | null): any;
