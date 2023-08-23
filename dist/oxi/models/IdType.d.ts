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
 *
 * @export
 * @interface IdType
 */
export interface IdType {
    /**
     * Used for Character Strings, length 0 to 20.
     * @type {string}
     * @memberof IdType
     */
    value?: string;
    /**
     *
     * @type {string}
     * @memberof IdType
     */
    source?: string;
}
/**
 * Check if a given object implements the IdType interface.
 */
export declare function instanceOfIdType(value: object): boolean;
export declare function IdTypeFromJSON(json: any): IdType;
export declare function IdTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdType;
export declare function IdTypeToJSON(value?: IdType | null): any;
