/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ErrorInstance } from './ErrorInstance';
import type { InstanceLink } from './InstanceLink';
/**
 * Complex type that contains error details for a REST call.
 * @export
 * @interface ExceptionDetailType
 */
export interface ExceptionDetailType {
    /**
     * Absolute URI [RFC3986] that identifies the problem type.  When dereferenced, it SHOULD provide a human-readable summary of the problem (for example, using HTML).
     * @type {string}
     * @memberof ExceptionDetailType
     */
    type: string;
    /**
     * Short, human-readable summary of the problem.  The summary SHOULD NOT change for subsequent occurrences of the problem, except for purposes of localization.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    title: string;
    /**
     * HTTP status code for this occurrence of the problem, set by the origin server.
     * @type {number}
     * @memberof ExceptionDetailType
     */
    status?: number;
    /**
     * Human-readable description specific to this occurrence of the problem.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    detail?: string;
    /**
     * Absolute URI that identifies the specific occurrence of the problem.  It may or may not provide additional information if dereferenced.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    instance?: string;
    /**
     * Application error code, which is different from HTTP error code.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    oerrorCode?: string;
    /**
     * Path to the problem at the resource or property level.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    oerrorPath?: string;
    /**
     * Details of the error message, consisting of a hierarchical tree structure.
     * @type {Array<ErrorInstance>}
     * @memberof ExceptionDetailType
     */
    oerrorDetails?: Array<ErrorInstance>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ExceptionDetailType
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the ExceptionDetailType interface.
 */
export declare function instanceOfExceptionDetailType(value: object): boolean;
export declare function ExceptionDetailTypeFromJSON(json: any): ExceptionDetailType;
export declare function ExceptionDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExceptionDetailType;
export declare function ExceptionDetailTypeToJSON(value?: ExceptionDetailType | null): any;
