/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AccumulatedBusinessEventType } from './AccumulatedBusinessEventType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface FetchAccumulatedBusinessEvents
 */
export interface FetchAccumulatedBusinessEvents {
    /**
     * List of accumulated business event messages
     * @type {Array<AccumulatedBusinessEventType>}
     * @memberof FetchAccumulatedBusinessEvents
     */
    businessEvents?: Array<AccumulatedBusinessEventType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof FetchAccumulatedBusinessEvents
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof FetchAccumulatedBusinessEvents
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof FetchAccumulatedBusinessEvents
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof FetchAccumulatedBusinessEvents
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof FetchAccumulatedBusinessEvents
     */
    totalResults?: number;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof FetchAccumulatedBusinessEvents
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FetchAccumulatedBusinessEvents
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the FetchAccumulatedBusinessEvents interface.
 */
export declare function instanceOfFetchAccumulatedBusinessEvents(value: object): boolean;
export declare function FetchAccumulatedBusinessEventsFromJSON(json: any): FetchAccumulatedBusinessEvents;
export declare function FetchAccumulatedBusinessEventsFromJSONTyped(json: any, ignoreDiscriminator: boolean): FetchAccumulatedBusinessEvents;
export declare function FetchAccumulatedBusinessEventsToJSON(value?: FetchAccumulatedBusinessEvents | null): any;
