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
import type { DateTimeSpanType } from './DateTimeSpanType';
/**
 *
 * @export
 * @interface TraceTimeInfoType
 */
export interface TraceTimeInfoType {
    /**
     *
     * @type {DateTimeSpanType}
     * @memberof TraceTimeInfoType
     */
    dateTimeSpan?: DateTimeSpanType;
    /**
     * Date of the trace.
     * @type {Date}
     * @memberof TraceTimeInfoType
     */
    traceOn?: Date;
    /**
     * Time of the trace
     * @type {string}
     * @memberof TraceTimeInfoType
     */
    traceTime?: string;
    /**
     * User that entered this trace.
     * @type {string}
     * @memberof TraceTimeInfoType
     */
    enteredBy?: string;
}
/**
 * Check if a given object implements the TraceTimeInfoType interface.
 */
export declare function instanceOfTraceTimeInfoType(value: object): boolean;
export declare function TraceTimeInfoTypeFromJSON(json: any): TraceTimeInfoType;
export declare function TraceTimeInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TraceTimeInfoType;
export declare function TraceTimeInfoTypeToJSON(value?: TraceTimeInfoType | null): any;
