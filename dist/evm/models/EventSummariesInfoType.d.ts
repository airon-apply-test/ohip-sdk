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
import type { EventSummaryInfoType } from './EventSummaryInfoType';
/**
 * Collection of events.
 * @export
 * @interface EventSummariesInfoType
 */
export interface EventSummariesInfoType {
    /**
     * Contains master/sub-events or standalone events summary information
     * @type {Array<EventSummaryInfoType>}
     * @memberof EventSummariesInfoType
     */
    eventSummary?: Array<EventSummaryInfoType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof EventSummariesInfoType
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof EventSummariesInfoType
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof EventSummariesInfoType
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof EventSummariesInfoType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof EventSummariesInfoType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof EventSummariesInfoType
     */
    count?: number;
}
/**
 * Check if a given object implements the EventSummariesInfoType interface.
 */
export declare function instanceOfEventSummariesInfoType(value: object): boolean;
export declare function EventSummariesInfoTypeFromJSON(json: any): EventSummariesInfoType;
export declare function EventSummariesInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EventSummariesInfoType;
export declare function EventSummariesInfoTypeToJSON(value?: EventSummariesInfoType | null): any;
