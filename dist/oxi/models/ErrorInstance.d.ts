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
/**
 * Complex type that contains error instance details for a REST call.
 * @export
 * @interface ErrorInstance
 */
export interface ErrorInstance {
    /**
     * Absolute URI [RFC3986] that identifies the problem type.  When dereferenced, it SHOULD provide a human-readable summary of the problem (for example, using HTML).
     * @type {string}
     * @memberof ErrorInstance
     */
    type: string;
    /**
     * Short, human-readable summary of the problem.  The summary SHOULD NOT change for subsequent occurrences of the problem, except for purposes of localization.
     * @type {string}
     * @memberof ErrorInstance
     */
    title: string;
    /**
     * HTTP status code for this occurrence of the problem, set by the origin server.
     * @type {number}
     * @memberof ErrorInstance
     */
    status?: number;
    /**
     * Human-readable description specific to this occurrence of the problem.
     * @type {string}
     * @memberof ErrorInstance
     */
    detail?: string;
    /**
     * Absolute URI that identifies the specific occurrence of the problem.  It may or may not provide additional information if dereferenced.
     * @type {string}
     * @memberof ErrorInstance
     */
    instance?: string;
    /**
     * Application error code, which is different from HTTP error code.
     * @type {string}
     * @memberof ErrorInstance
     */
    oerrorCode?: string;
    /**
     * Path to the problem at the resource or property level.
     * @type {string}
     * @memberof ErrorInstance
     */
    oerrorPath?: string;
}
/**
 * Check if a given object implements the ErrorInstance interface.
 */
export declare function instanceOfErrorInstance(value: object): boolean;
export declare function ErrorInstanceFromJSON(json: any): ErrorInstance;
export declare function ErrorInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ErrorInstance;
export declare function ErrorInstanceToJSON(value?: ErrorInstance | null): any;
